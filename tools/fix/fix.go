// Copyright 2019 CUE Authors
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

// Package fix contains functionality for writing CUE files with legacy
// syntax to newer ones.
//
// Note: the transformations that are supported in this package will change
// over time.
package fix

import (
	"fmt"
	"strings"

	"github.com/solo-io/cue/cue/ast"
	"github.com/solo-io/cue/cue/ast/astutil"
	"github.com/solo-io/cue/cue/token"
)

type Option func(*options)

type options struct {
	simplify   bool
	deprecated bool
}

// Simplify enables fixes that simplify the code, but are not strictly
// necessary.
func Simplify() Option {
	return func(o *options) { o.simplify = true }
}

// File applies fixes to f and returns it. It alters the original f.
func File(f *ast.File, o ...Option) *ast.File {
	var options options
	for _, f := range o {
		f(&options)
	}

	// Rewrite integer division operations to use builtins.
	f = astutil.Apply(f, func(c astutil.Cursor) bool {
		n := c.Node()
		switch x := n.(type) {
		case *ast.BinaryExpr:
			switch x.Op {
			case token.IDIV, token.IMOD, token.IQUO, token.IREM:
				ast.SetRelPos(x.X, token.NoSpace)
				c.Replace(&ast.CallExpr{
					// Use the __foo version to prevent accidental shadowing.
					Fun:  ast.NewIdent("__" + x.Op.String()),
					Args: []ast.Expr{x.X, x.Y},
				})
			}
		}
		return true
	}, nil).(*ast.File)

	// Rewrite an old-style alias to a let clause.
	ast.Walk(f, func(n ast.Node) bool {
		var decls []ast.Decl
		switch x := n.(type) {
		case *ast.StructLit:
			decls = x.Elts
		case *ast.File:
			decls = x.Decls
		}
		for i, d := range decls {
			if a, ok := d.(*ast.Alias); ok {
				x := &ast.LetClause{
					Ident: a.Ident,
					Equal: a.Equal,
					Expr:  a.Expr,
				}
				astutil.CopyMeta(x, a)
				decls[i] = x
			}
		}
		return true
	}, nil)

	// Rewrite block comments to regular comments.
	ast.Walk(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.CommentGroup:
			comments := []*ast.Comment{}
			for _, c := range x.List {
				s := c.Text
				if !strings.HasPrefix(s, "/*") || !strings.HasSuffix(s, "*/") {
					comments = append(comments, c)
					continue
				}
				if x.Position > 0 {
					// Moving to the end doesn't work, as it still
					// may inject at a false line break position.
					x.Position = 0
					x.Doc = true
				}
				s = strings.TrimSpace(s[2 : len(s)-2])
				for _, s := range strings.Split(s, "\n") {
					for i := 0; i < 3; i++ {
						if strings.HasPrefix(s, " ") || strings.HasPrefix(s, "*") {
							s = s[1:]
						}
					}
					comments = append(comments, &ast.Comment{Text: "// " + s})
				}
			}
			x.List = comments
			return false
		}
		return true
	}, nil)

	// Referred nodes and used identifiers.
	referred := map[ast.Node]string{}
	used := map[string]bool{}
	replacement := map[ast.Node]string{}

	ast.Walk(f, func(n ast.Node) bool {
		if i, ok := n.(*ast.Ident); ok {
			str, err := ast.ParseIdent(i)
			if err != nil {
				return false
			}
			referred[i.Node] = str
			used[str] = true
		}
		return true
	}, nil)

	num := 0
	newIdent := func() string {
		for num++; ; num++ {
			str := fmt.Sprintf("X%d", num)
			if !used[str] {
				used[str] = true
				return str
			}
		}
	}

	// Rewrite quoted identifier fields that are referenced.
	f = astutil.Apply(f, func(c astutil.Cursor) bool {
		n := c.Node()
		switch x := n.(type) {
		case *ast.Field:
			m, ok := referred[x.Value]
			if !ok {
				break
			}

			if b, ok := x.Label.(*ast.Ident); ok {
				str, err := ast.ParseIdent(b)
				var expr ast.Expr = b

				switch {
				case token.Lookup(str) != token.IDENT:
					// quote keywords
					expr = ast.NewString(b.Name)

				case err != nil || str != m || str == b.Name:
					return true

				case ast.IsValidIdent(str):
					x.Label = astutil.CopyMeta(ast.NewIdent(str), x.Label).(ast.Label)
					return true
				}

				ident := newIdent()
				replacement[x.Value] = ident
				expr = &ast.Alias{Ident: ast.NewIdent(ident), Expr: expr}
				ast.SetRelPos(x.Label, token.NoRelPos)
				x.Label = astutil.CopyMeta(expr, x.Label).(ast.Label)
			}
		}
		return true
	}, nil).(*ast.File)

	// Replace quoted references with their alias identifier.
	astutil.Apply(f, func(c astutil.Cursor) bool {
		n := c.Node()
		switch x := n.(type) {
		case *ast.Ident:
			if r, ok := replacement[x.Node]; ok {
				c.Replace(astutil.CopyMeta(ast.NewIdent(r), n))
				break
			}
			str, err := ast.ParseIdent(x)
			if err != nil || str == x.Name {
				break
			}
			// Either the identifier is valid, in which can be replaced simply
			// as here, or it is a complicated identifier and the original
			// destination must have been quoted, in which case it is handled
			// above.
			if ast.IsValidIdent(str) && token.Lookup(str) == token.IDENT {
				c.Replace(astutil.CopyMeta(ast.NewIdent(str), n))
			}
		}
		return true
	}, nil)

	// TODO: we are probably reintroducing slices. Disable for now.
	//
	// Rewrite slice expression.
	// f = astutil.Apply(f, func(c astutil.Cursor) bool {
	// 	n := c.Node()
	// 	getVal := func(n ast.Expr) ast.Expr {
	// 		if n == nil {
	// 			return nil
	// 		}
	// 		if id, ok := n.(*ast.Ident); ok && id.Name == "_" {
	// 			return nil
	// 		}
	// 		return n
	// 	}
	// 	switch x := n.(type) {
	// 	case *ast.SliceExpr:
	// 		ast.SetRelPos(x.X, token.NoRelPos)

	// 		lo := getVal(x.Low)
	// 		hi := getVal(x.High)
	// 		if lo == nil { // a[:j]
	// 			lo = mustParseExpr("0")
	// 			astutil.CopyMeta(lo, x.Low)
	// 		}
	// 		if hi == nil { // a[i:]
	// 			hi = ast.NewCall(ast.NewIdent("len"), x.X)
	// 			astutil.CopyMeta(lo, x.High)
	// 		}
	// 		if pkg := c.Import("list"); pkg != nil {
	// 			c.Replace(ast.NewCall(ast.NewSel(pkg, "Slice"), x.X, lo, hi))
	// 		}
	// 	}
	// 	return true
	// }, nil).(*ast.File)

	if options.simplify {
		f = simplify(f)
	}

	return f
}
