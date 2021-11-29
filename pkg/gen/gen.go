// Copyright 2018 The CUE Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package gen is a command that can be used to bootstrap a new builtin package
// directory. The directory has to reside in cuelang.org/go/pkg.
//
// To bootstrap a directory, run this command from within that direcory.
// After that directory's files can be regenerated with go generate.
//
// Be sure to also update an entry in pkg/pkg.go, if so desired.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/constant"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/solo-io/cue/cue"
	"github.com/solo-io/cue/cue/errors"
	cueformat "github.com/solo-io/cue/cue/format"
	"github.com/solo-io/cue/cue/load"
	"github.com/solo-io/cue/internal"
)

const genFile = "pkg.go"

const prefix = "../pkg/"

const header = `// Code generated by go generate. DO NOT EDIT.

 //go:generate rm %s
 //go:generate go run %sgen/gen.go

package %s

import (
	"github.com/solo-io/cue/internal/core/adt"
	"github.com/solo-io/cue/pkg/internal"
)

func init() {
	internal.Register(%q, pkg)
}

var _ = adt.TopKind // in case the adt package isn't used

`

func main() {
	flag.Parse()
	log.SetFlags(log.Lshortfile)
	log.SetOutput(os.Stdout)

	g := generator{
		w:     &bytes.Buffer{},
		decls: &bytes.Buffer{},
		fset:  token.NewFileSet(),
	}

	cwd, _ := os.Getwd()
	pkg := strings.Split(filepath.ToSlash(cwd), "/pkg/")[1]
	gopkg := path.Base(pkg)
	// TODO: rename list to lists and struct to structs.
	if gopkg == "struct" {
		gopkg = "structs"
	}
	dots := strings.Repeat("../", strings.Count(pkg, "/")+1)

	w := &bytes.Buffer{}
	fmt.Fprintf(w, header, genFile, dots, gopkg, pkg)
	g.processDir(pkg)

	io.Copy(w, g.decls)
	io.Copy(w, g.w)

	b, err := format.Source(w.Bytes())
	if err != nil {
		b = w.Bytes() // write the unformatted source
	}

	b = bytes.Replace(b, []byte(".Builtin{{}}"), []byte(".Builtin{}"), -1)

	filename := filepath.Join(genFile)

	if err := ioutil.WriteFile(filename, b, 0666); err != nil {
		log.Fatal(err)
	}
}

type generator struct {
	w          *bytes.Buffer
	decls      *bytes.Buffer
	name       string
	fset       *token.FileSet
	defaultPkg string
	first      bool
	iota       int

	imports []*ast.ImportSpec
}

func (g *generator) processDir(pkg string) {
	goFiles, err := filepath.Glob("*.go")
	if err != nil {
		log.Fatal(err)
	}

	cueFiles, err := filepath.Glob("*.cue")
	if err != nil {
		log.Fatal(err)
	}

	if len(goFiles)+len(cueFiles) == 0 {
		return
	}

	fmt.Fprintf(g.w, "var pkg = &internal.Package{\nNative: []*internal.Builtin{{\n")
	g.first = true
	for _, filename := range goFiles {
		if filename == genFile {
			continue
		}
		g.processGo(filename)
	}
	fmt.Fprintf(g.w, "}},\n")
	g.processCUE(pkg)
	fmt.Fprintf(g.w, "}\n")
}

func (g *generator) sep() {
	if g.first {
		g.first = false
		return
	}
	fmt.Fprintln(g.w, "}, {")
}

// processCUE mixes in CUE definitions defined in the package directory.
func (g *generator) processCUE(pkg string) {
	instances := cue.Build(load.Instances([]string{"."}, &load.Config{
		StdRoot: ".",
	}))

	if err := instances[0].Err; err != nil {
		if !strings.Contains(err.Error(), "no CUE files") {
			errors.Print(os.Stderr, err, nil)
			log.Fatalf("error processing %s: %v", pkg, err)
		}
		return
	}

	v := instances[0].Value().Syntax(cue.Raw())
	// fmt.Printf("%T\n", v)
	// fmt.Println(astinternal.DebugStr(v))
	n := internal.ToExpr(v)
	b, err := cueformat.Node(n)
	if err != nil {
		log.Fatal(err)
	}
	b = bytes.ReplaceAll(b, []byte("\n\n"), []byte("\n"))
	// body = strings.ReplaceAll(body, "\t", "")
	// TODO: escape backtick
	fmt.Fprintf(g.w, "CUE: `%s`,\n", string(b))
}

func (g *generator) processGo(filename string) {
	if strings.HasSuffix(filename, "_test.go") {
		return
	}
	f, err := parser.ParseFile(g.fset, filename, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	g.defaultPkg = ""
	g.name = f.Name.Name
	if g.name == "structs" {
		g.name = "struct"
	}

	for _, d := range f.Decls {
		switch x := d.(type) {
		case *ast.GenDecl:
			switch x.Tok {
			case token.CONST:
				for _, spec := range x.Specs {
					if !ast.IsExported(spec.(*ast.ValueSpec).Names[0].Name) {
						continue
					}
					g.genConst(spec.(*ast.ValueSpec))
				}
			case token.VAR:
				continue
			case token.TYPE:
				// TODO: support type declarations.
				continue
			case token.IMPORT:
				continue
			default:
				log.Fatalf("gen %s: unexpected spec of type %s", filename, x.Tok)
			}
		case *ast.FuncDecl:
			g.genFun(x)
		}
	}
}

func (g *generator) genConst(spec *ast.ValueSpec) {
	name := spec.Names[0].Name
	value := ""
	switch v := g.toValue(spec.Values[0]); v.Kind() {
	case constant.Bool, constant.Int, constant.String:
		// TODO: convert octal numbers
		value = v.ExactString()
	case constant.Float:
		var rat big.Rat
		rat.SetString(v.ExactString())
		var float big.Float
		float.SetRat(&rat)
		value = float.Text('g', -1)
	default:
		fmt.Printf("Dropped entry %s.%s (%T: %v)\n", g.defaultPkg, name, v.Kind(), v.ExactString())
		return
	}
	g.sep()
	fmt.Fprintf(g.w, "Name: %q,\n Const: %q,\n", name, value)
}

func (g *generator) toValue(x ast.Expr) constant.Value {
	switch x := x.(type) {
	case *ast.BasicLit:
		return constant.MakeFromLiteral(x.Value, x.Kind, 0)
	case *ast.BinaryExpr:
		return constant.BinaryOp(g.toValue(x.X), x.Op, g.toValue(x.Y))
	case *ast.UnaryExpr:
		return constant.UnaryOp(x.Op, g.toValue(x.X), 0)
	default:
		log.Fatalf("%s: unsupported expression type %T: %#v", g.defaultPkg, x, x)
	}
	return constant.MakeUnknown()
}

func (g *generator) genFun(x *ast.FuncDecl) {
	if x.Body == nil || !ast.IsExported(x.Name.Name) {
		return
	}
	types := []string{}
	if x.Type.Results != nil {
		for _, f := range x.Type.Results.List {
			if len(f.Names) > 0 {
				for range f.Names {
					types = append(types, g.goKind(f.Type))
				}
			} else {
				types = append(types, g.goKind(f.Type))
			}
		}
	}
	if n := len(types); n != 1 && (n != 2 || types[1] != "error") {
		fmt.Printf("Dropped func %s.%s: must have one return value or a value and an error %v\n", g.defaultPkg, x.Name.Name, types)
		return
	}

	if x.Recv != nil {
		// if strings.HasPrefix(x.Name.Name, g.name) {
		// 	printer.Fprint(g.decls, g.fset, x)
		// 	fmt.Fprint(g.decls, "\n\n")
		// }
		return
	}

	g.sep()
	fmt.Fprintf(g.w, "Name: %q,\n", x.Name.Name)

	args := []string{}
	vals := []string{}
	kind := []string{}
	for _, f := range x.Type.Params.List {
		for _, name := range f.Names {
			typ := strings.Title(g.goKind(f.Type))
			argKind := g.goToCUE(f.Type)
			vals = append(vals, fmt.Sprintf("c.%s(%d)", typ, len(args)))
			args = append(args, name.Name)
			kind = append(kind, argKind)
		}
	}

	fmt.Fprintf(g.w, "Params: []internal.Param{\n")
	for _, k := range kind {
		fmt.Fprintf(g.w, "{Kind: %s},\n", k)
	}
	fmt.Fprintf(g.w, "\n},\n")

	expr := x.Type.Results.List[0].Type
	fmt.Fprintf(g.w, "Result: %s,\n", g.goToCUE(expr))

	argList := strings.Join(args, ", ")
	valList := strings.Join(vals, ", ")
	init := ""
	if len(args) > 0 {
		init = fmt.Sprintf("%s := %s", argList, valList)
	}

	fmt.Fprintf(g.w, "Func: func(c *internal.CallCtxt) {")
	defer fmt.Fprintln(g.w, "},")
	fmt.Fprintln(g.w)
	if init != "" {
		fmt.Fprintln(g.w, init)
	}
	fmt.Fprintln(g.w, "if c.Do() {")
	defer fmt.Fprintln(g.w, "}")
	if len(types) == 1 {
		fmt.Fprintf(g.w, "c.Ret = %s(%s)", x.Name.Name, argList)
	} else {
		fmt.Fprintf(g.w, "c.Ret, c.Err = %s(%s)", x.Name.Name, argList)
	}
}

func (g *generator) goKind(expr ast.Expr) string {
	if star, isStar := expr.(*ast.StarExpr); isStar {
		expr = star.X
	}
	w := &bytes.Buffer{}
	printer.Fprint(w, g.fset, expr)
	switch str := w.String(); str {
	case "big.Int":
		return "bigInt"
	case "big.Float":
		return "bigFloat"
	case "big.Rat":
		return "bigRat"
	case "adt.Bottom":
		return "error"
	case "internal.Decimal":
		return "decimal"
	case "[]*internal.Decimal":
		return "decimalList"
	case "cue.Struct":
		return "struct"
	case "cue.Value":
		return "value"
	case "cue.List":
		return "list"
	case "[]string":
		return "stringList"
	case "[]byte":
		return "bytes"
	case "[]cue.Value":
		return "list"
	case "io.Reader":
		return "reader"
	case "time.Time":
		return "string"
	default:
		return str
	}
}

func (g *generator) goToCUE(expr ast.Expr) (cueKind string) {
	// TODO: detect list and structs types for return values.
	switch k := g.goKind(expr); k {
	case "error":
		cueKind += "adt.BottomKind"
	case "bool":
		cueKind += "adt.BoolKind"
	case "bytes", "reader":
		cueKind += "adt.BytesKind|adt.StringKind"
	case "string":
		cueKind += "adt.StringKind"
	case "int", "int8", "int16", "int32", "rune", "int64",
		"uint", "byte", "uint8", "uint16", "uint32", "uint64",
		"bigInt":
		cueKind += "adt.IntKind"
	case "float64", "bigRat", "bigFloat", "decimal":
		cueKind += "adt.NumKind"
	case "list", "decimalList", "stringList":
		cueKind += "adt.ListKind"
	case "struct":
		cueKind += "adt.StructKind"
	case "value":
		// Must use callCtxt.value method for these types and resolve manually.
		cueKind += "adt.TopKind" // TODO: can be more precise
	default:
		switch {
		case strings.HasPrefix(k, "[]"):
			cueKind += "adt.ListKind"
		case strings.HasPrefix(k, "map["):
			cueKind += "adt.StructKind"
		default:
			// log.Println("Unknown type:", k)
			// Must use callCtxt.value method for these types and resolve manually.
			cueKind += "adt.TopKind" // TODO: can be more precise
		}
	}
	return cueKind
}
