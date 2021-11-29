// Code generated by go generate. DO NOT EDIT.

//go:generate rm pkg.go
//go:generate go run ../gen/gen.go

package uuid

import (
	"github.com/solo-io/cue/internal/core/adt"
	"github.com/solo-io/cue/pkg/internal"
)

func init() {
	internal.Register("uuid", pkg)
}

var _ = adt.TopKind // in case the adt package isn't used

var pkg = &internal.Package{
	Native: []*internal.Builtin{{
		Name: "Valid",
		Params: []internal.Param{
			{Kind: adt.StringKind},
		},
		Result: adt.BottomKind,
		Func: func(c *internal.CallCtxt) {
			s := c.String(0)
			if c.Do() {
				c.Ret = Valid(s)
			}
		},
	}, {
		Name: "Parse",
		Params: []internal.Param{
			{Kind: adt.StringKind},
		},
		Result: adt.StringKind,
		Func: func(c *internal.CallCtxt) {
			s := c.String(0)
			if c.Do() {
				c.Ret, c.Err = Parse(s)
			}
		},
	}, {
		Name: "ToString",
		Params: []internal.Param{
			{Kind: adt.StringKind},
		},
		Result: adt.StringKind,
		Func: func(c *internal.CallCtxt) {
			x := c.String(0)
			if c.Do() {
				c.Ret = ToString(x)
			}
		},
	}, {
		Name: "URN",
		Params: []internal.Param{
			{Kind: adt.StringKind},
		},
		Result: adt.StringKind,
		Func: func(c *internal.CallCtxt) {
			x := c.String(0)
			if c.Do() {
				c.Ret, c.Err = URN(x)
			}
		},
	}, {
		Name: "FromInt",
		Params: []internal.Param{
			{Kind: adt.IntKind},
		},
		Result: adt.StringKind,
		Func: func(c *internal.CallCtxt) {
			i := c.BigInt(0)
			if c.Do() {
				c.Ret, c.Err = FromInt(i)
			}
		},
	}, {
		Name: "ToInt",
		Params: []internal.Param{
			{Kind: adt.StringKind},
		},
		Result: adt.IntKind,
		Func: func(c *internal.CallCtxt) {
			x := c.String(0)
			if c.Do() {
				c.Ret = ToInt(x)
			}
		},
	}, {
		Name: "Variant",
		Params: []internal.Param{
			{Kind: adt.StringKind},
		},
		Result: adt.IntKind,
		Func: func(c *internal.CallCtxt) {
			x := c.String(0)
			if c.Do() {
				c.Ret, c.Err = Variant(x)
			}
		},
	}, {
		Name: "Version",
		Params: []internal.Param{
			{Kind: adt.StringKind},
		},
		Result: adt.IntKind,
		Func: func(c *internal.CallCtxt) {
			x := c.String(0)
			if c.Do() {
				c.Ret, c.Err = Version(x)
			}
		},
	}, {
		Name: "SHA1",
		Params: []internal.Param{
			{Kind: adt.StringKind},
			{Kind: adt.BytesKind | adt.StringKind},
		},
		Result: adt.StringKind,
		Func: func(c *internal.CallCtxt) {
			space, data := c.String(0), c.Bytes(1)
			if c.Do() {
				c.Ret, c.Err = SHA1(space, data)
			}
		},
	}, {
		Name: "MD5",
		Params: []internal.Param{
			{Kind: adt.StringKind},
			{Kind: adt.BytesKind | adt.StringKind},
		},
		Result: adt.StringKind,
		Func: func(c *internal.CallCtxt) {
			space, data := c.String(0), c.Bytes(1)
			if c.Do() {
				c.Ret, c.Err = MD5(space, data)
			}
		},
	}},
	CUE: `{
	ns: {
		DNS:  "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
		URL:  "6ba7b811-9dad-11d1-80b4-00c04fd430c8"
		OID:  "6ba7b812-9dad-11d1-80b4-00c04fd430c8"
		X500: "6ba7b814-9dad-11d1-80b4-00c04fd430c8"
		Nil:  "00000000-0000-0000-0000-000000000000"
	}
	variants: {
		Invalid:   0
		RFC4122:   1
		Reserved:  2
		Microsoft: 3
		Future:    4
	}
}`,
}
