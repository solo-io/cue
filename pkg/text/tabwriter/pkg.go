// Code generated by go generate. DO NOT EDIT.

//go:generate rm pkg.go
//go:generate go run ../../gen/gen.go

package tabwriter

import (
	"github.com/solo-io/cue/internal/core/adt"
	"github.com/solo-io/cue/pkg/internal"
)

func init() {
	internal.Register("text/tabwriter", pkg)
}

var _ = adt.TopKind // in case the adt package isn't used

var pkg = &internal.Package{
	Native: []*internal.Builtin{{
		Name: "Write",
		Params: []internal.Param{
			{Kind: adt.TopKind},
		},
		Result: adt.StringKind,
		Func: func(c *internal.CallCtxt) {
			data := c.Value(0)
			if c.Do() {
				c.Ret, c.Err = Write(data)
			}
		},
	}},
}
