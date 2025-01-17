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

package load

import (
	"os"
	"reflect"
	"testing"

	"github.com/solo-io/cue/cue/build"
	"github.com/solo-io/cue/cue/token"
)

const testdata = "./testdata/"

func getInst(pkg, cwd string) (*build.Instance, error) {
	c, _ := (&Config{Dir: cwd}).complete()
	l := loader{cfg: c}
	inst := c.newRelInstance(token.NoPos, pkg, c.Package)
	p := l.importPkg(token.NoPos, inst)[0]
	return p, p.Err
}

func TestEmptyImport(t *testing.T) {
	p, err := getInst("", "")
	if err == nil {
		t.Fatal(`Import("") returned nil error.`)
	}
	if p == nil {
		t.Fatal(`Import("") returned nil package.`)
	}
	if p.DisplayPath != "" {
		t.Fatalf("DisplayPath=%q, want %q.", p.DisplayPath, "")
	}
}

func TestEmptyFolderImport(t *testing.T) {
	_, err := getInst(".", testdata+"empty")
	if _, ok := err.(*NoFilesError); !ok {
		t.Fatal(`Import("testdata/empty") did not return NoCUEError.`)
	}
}

func TestMultiplePackageImport(t *testing.T) {
	_, err := getInst(".", testdata+"multi")
	mpe, ok := err.(*MultiplePackageError)
	if !ok {
		t.Fatal(`Import("testdata/multi") did not return MultiplePackageError.`)
	}
	mpe.Dir = ""
	want := &MultiplePackageError{
		Packages: []string{"main", "test_package"},
		Files:    []string{"file.cue", "file_appengine.cue"},
	}
	if !reflect.DeepEqual(mpe, want) {
		t.Errorf("got %#v; want %#v", mpe, want)
	}
}

func TestLocalDirectory(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	p, err := getInst(".", cwd)
	if err != nil {
		t.Fatal(err)
	}

	if p.DisplayPath != "." {
		t.Fatalf("DisplayPath=%q, want %q", p.DisplayPath, ".")
	}
}
