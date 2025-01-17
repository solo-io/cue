// Code generated by go generate. DO NOT EDIT.

//go:generate rm pkg.go
//go:generate go run ../gen/gen.go

package regexp

import (
	"github.com/solo-io/cue/internal/core/adt"
	"github.com/solo-io/cue/pkg/internal"
)

func init() {
	internal.Register("regexp", pkg)
}

var _ = adt.TopKind // in case the adt package isn't used

var pkg = &internal.Package{
	Native: []*internal.Builtin{{
		Name: "Valid",
		Params: []internal.Param{
			{Kind: adt.StringKind},
		},
		Result: adt.BoolKind,
		Func: func(c *internal.CallCtxt) {
			pattern := c.String(0)
			if c.Do() {
				c.Ret, c.Err = Valid(pattern)
			}
		},
	}, {
		Name: "Find",
		Params: []internal.Param{
			{Kind: adt.StringKind},
			{Kind: adt.StringKind},
		},
		Result: adt.StringKind,
		Func: func(c *internal.CallCtxt) {
			pattern, s := c.String(0), c.String(1)
			if c.Do() {
				c.Ret, c.Err = Find(pattern, s)
			}
		},
	}, {
		Name: "FindAll",
		Params: []internal.Param{
			{Kind: adt.StringKind},
			{Kind: adt.StringKind},
			{Kind: adt.IntKind},
		},
		Result: adt.ListKind,
		Func: func(c *internal.CallCtxt) {
			pattern, s, n := c.String(0), c.String(1), c.Int(2)
			if c.Do() {
				c.Ret, c.Err = FindAll(pattern, s, n)
			}
		},
	}, {
		Name: "FindSubmatch",
		Params: []internal.Param{
			{Kind: adt.StringKind},
			{Kind: adt.StringKind},
		},
		Result: adt.ListKind,
		Func: func(c *internal.CallCtxt) {
			pattern, s := c.String(0), c.String(1)
			if c.Do() {
				c.Ret, c.Err = FindSubmatch(pattern, s)
			}
		},
	}, {
		Name: "FindAllSubmatch",
		Params: []internal.Param{
			{Kind: adt.StringKind},
			{Kind: adt.StringKind},
			{Kind: adt.IntKind},
		},
		Result: adt.ListKind,
		Func: func(c *internal.CallCtxt) {
			pattern, s, n := c.String(0), c.String(1), c.Int(2)
			if c.Do() {
				c.Ret, c.Err = FindAllSubmatch(pattern, s, n)
			}
		},
	}, {
		Name: "FindNamedSubmatch",
		Params: []internal.Param{
			{Kind: adt.StringKind},
			{Kind: adt.StringKind},
		},
		Result: adt.StructKind,
		Func: func(c *internal.CallCtxt) {
			pattern, s := c.String(0), c.String(1)
			if c.Do() {
				c.Ret, c.Err = FindNamedSubmatch(pattern, s)
			}
		},
	}, {
		Name: "FindAllNamedSubmatch",
		Params: []internal.Param{
			{Kind: adt.StringKind},
			{Kind: adt.StringKind},
			{Kind: adt.IntKind},
		},
		Result: adt.ListKind,
		Func: func(c *internal.CallCtxt) {
			pattern, s, n := c.String(0), c.String(1), c.Int(2)
			if c.Do() {
				c.Ret, c.Err = FindAllNamedSubmatch(pattern, s, n)
			}
		},
	}, {
		Name: "Match",
		Params: []internal.Param{
			{Kind: adt.StringKind},
			{Kind: adt.StringKind},
		},
		Result: adt.BoolKind,
		Func: func(c *internal.CallCtxt) {
			pattern, s := c.String(0), c.String(1)
			if c.Do() {
				c.Ret, c.Err = Match(pattern, s)
			}
		},
	}, {
		Name: "QuoteMeta",
		Params: []internal.Param{
			{Kind: adt.StringKind},
		},
		Result: adt.StringKind,
		Func: func(c *internal.CallCtxt) {
			s := c.String(0)
			if c.Do() {
				c.Ret = QuoteMeta(s)
			}
		},
	}},
}
