// Code generated by go generate. DO NOT EDIT.

//go:generate rm pkg.go
//go:generate go run ../../gen/gen.go

package md5

import (
	"github.com/solo-io/cue/internal/core/adt"
	"github.com/solo-io/cue/pkg/internal"
)

func init() {
	internal.Register("crypto/md5", pkg)
}

var _ = adt.TopKind // in case the adt package isn't used

var pkg = &internal.Package{
	Native: []*internal.Builtin{{
		Name:  "Size",
		Const: "16",
	}, {
		Name:  "BlockSize",
		Const: "64",
	}, {
		Name: "Sum",
		Params: []internal.Param{
			{Kind: adt.BytesKind | adt.StringKind},
		},
		Result: adt.BytesKind | adt.StringKind,
		Func: func(c *internal.CallCtxt) {
			data := c.Bytes(0)
			if c.Do() {
				c.Ret = Sum(data)
			}
		},
	}},
}
