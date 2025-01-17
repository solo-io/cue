// Code generated by go generate. DO NOT EDIT.

//go:generate rm pkg.go
//go:generate go run ../gen/gen.go

package list

import (
	"github.com/solo-io/cue/internal/core/adt"
	"github.com/solo-io/cue/pkg/internal"
)

func init() {
	internal.Register("list", pkg)
}

var _ = adt.TopKind // in case the adt package isn't used

var pkg = &internal.Package{
	Native: []*internal.Builtin{{
		Name: "Drop",
		Params: []internal.Param{
			{Kind: adt.ListKind},
			{Kind: adt.IntKind},
		},
		Result: adt.ListKind,
		Func: func(c *internal.CallCtxt) {
			x, n := c.List(0), c.Int(1)
			if c.Do() {
				c.Ret, c.Err = Drop(x, n)
			}
		},
	}, {
		Name: "FlattenN",
		Params: []internal.Param{
			{Kind: adt.TopKind},
			{Kind: adt.IntKind},
		},
		Result: adt.ListKind,
		Func: func(c *internal.CallCtxt) {
			xs, depth := c.Value(0), c.Int(1)
			if c.Do() {
				c.Ret, c.Err = FlattenN(xs, depth)
			}
		},
	}, {
		Name: "Repeat",
		Params: []internal.Param{
			{Kind: adt.ListKind},
			{Kind: adt.IntKind},
		},
		Result: adt.ListKind,
		Func: func(c *internal.CallCtxt) {
			x, count := c.List(0), c.Int(1)
			if c.Do() {
				c.Ret, c.Err = Repeat(x, count)
			}
		},
	}, {
		Name: "Concat",
		Params: []internal.Param{
			{Kind: adt.ListKind},
		},
		Result: adt.ListKind,
		Func: func(c *internal.CallCtxt) {
			a := c.List(0)
			if c.Do() {
				c.Ret, c.Err = Concat(a)
			}
		},
	}, {
		Name: "Take",
		Params: []internal.Param{
			{Kind: adt.ListKind},
			{Kind: adt.IntKind},
		},
		Result: adt.ListKind,
		Func: func(c *internal.CallCtxt) {
			x, n := c.List(0), c.Int(1)
			if c.Do() {
				c.Ret, c.Err = Take(x, n)
			}
		},
	}, {
		Name: "Slice",
		Params: []internal.Param{
			{Kind: adt.ListKind},
			{Kind: adt.IntKind},
			{Kind: adt.IntKind},
		},
		Result: adt.ListKind,
		Func: func(c *internal.CallCtxt) {
			x, i, j := c.List(0), c.Int(1), c.Int(2)
			if c.Do() {
				c.Ret, c.Err = Slice(x, i, j)
			}
		},
	}, {
		Name: "MinItems",
		Params: []internal.Param{
			{Kind: adt.ListKind},
			{Kind: adt.IntKind},
		},
		Result: adt.BoolKind,
		Func: func(c *internal.CallCtxt) {
			a, n := c.List(0), c.Int(1)
			if c.Do() {
				c.Ret = MinItems(a, n)
			}
		},
	}, {
		Name: "MaxItems",
		Params: []internal.Param{
			{Kind: adt.ListKind},
			{Kind: adt.IntKind},
		},
		Result: adt.BoolKind,
		Func: func(c *internal.CallCtxt) {
			a, n := c.List(0), c.Int(1)
			if c.Do() {
				c.Ret = MaxItems(a, n)
			}
		},
	}, {
		Name: "UniqueItems",
		Params: []internal.Param{
			{Kind: adt.ListKind},
		},
		Result: adt.BoolKind,
		Func: func(c *internal.CallCtxt) {
			a := c.List(0)
			if c.Do() {
				c.Ret = UniqueItems(a)
			}
		},
	}, {
		Name: "Contains",
		Params: []internal.Param{
			{Kind: adt.ListKind},
			{Kind: adt.TopKind},
		},
		Result: adt.BoolKind,
		Func: func(c *internal.CallCtxt) {
			a, v := c.List(0), c.Value(1)
			if c.Do() {
				c.Ret = Contains(a, v)
			}
		},
	}, {
		Name: "Avg",
		Params: []internal.Param{
			{Kind: adt.ListKind},
		},
		Result: adt.NumKind,
		Func: func(c *internal.CallCtxt) {
			xs := c.DecimalList(0)
			if c.Do() {
				c.Ret, c.Err = Avg(xs)
			}
		},
	}, {
		Name: "Max",
		Params: []internal.Param{
			{Kind: adt.ListKind},
		},
		Result: adt.NumKind,
		Func: func(c *internal.CallCtxt) {
			xs := c.DecimalList(0)
			if c.Do() {
				c.Ret, c.Err = Max(xs)
			}
		},
	}, {
		Name: "Min",
		Params: []internal.Param{
			{Kind: adt.ListKind},
		},
		Result: adt.NumKind,
		Func: func(c *internal.CallCtxt) {
			xs := c.DecimalList(0)
			if c.Do() {
				c.Ret, c.Err = Min(xs)
			}
		},
	}, {
		Name: "Product",
		Params: []internal.Param{
			{Kind: adt.ListKind},
		},
		Result: adt.NumKind,
		Func: func(c *internal.CallCtxt) {
			xs := c.DecimalList(0)
			if c.Do() {
				c.Ret, c.Err = Product(xs)
			}
		},
	}, {
		Name: "Range",
		Params: []internal.Param{
			{Kind: adt.NumKind},
			{Kind: adt.NumKind},
			{Kind: adt.NumKind},
		},
		Result: adt.ListKind,
		Func: func(c *internal.CallCtxt) {
			start, limit, step := c.Decimal(0), c.Decimal(1), c.Decimal(2)
			if c.Do() {
				c.Ret, c.Err = Range(start, limit, step)
			}
		},
	}, {
		Name: "Sum",
		Params: []internal.Param{
			{Kind: adt.ListKind},
		},
		Result: adt.NumKind,
		Func: func(c *internal.CallCtxt) {
			xs := c.DecimalList(0)
			if c.Do() {
				c.Ret, c.Err = Sum(xs)
			}
		},
	}, {
		Name: "Sort",
		Params: []internal.Param{
			{Kind: adt.ListKind},
			{Kind: adt.TopKind},
		},
		Result: adt.ListKind,
		Func: func(c *internal.CallCtxt) {
			list, cmp := c.List(0), c.Value(1)
			if c.Do() {
				c.Ret, c.Err = Sort(list, cmp)
			}
		},
	}, {
		Name: "SortStable",
		Params: []internal.Param{
			{Kind: adt.ListKind},
			{Kind: adt.TopKind},
		},
		Result: adt.ListKind,
		Func: func(c *internal.CallCtxt) {
			list, cmp := c.List(0), c.Value(1)
			if c.Do() {
				c.Ret, c.Err = SortStable(list, cmp)
			}
		},
	}, {
		Name: "SortStrings",
		Params: []internal.Param{
			{Kind: adt.ListKind},
		},
		Result: adt.ListKind,
		Func: func(c *internal.CallCtxt) {
			a := c.StringList(0)
			if c.Do() {
				c.Ret = SortStrings(a)
			}
		},
	}, {
		Name: "IsSorted",
		Params: []internal.Param{
			{Kind: adt.ListKind},
			{Kind: adt.TopKind},
		},
		Result: adt.BoolKind,
		Func: func(c *internal.CallCtxt) {
			list, cmp := c.List(0), c.Value(1)
			if c.Do() {
				c.Ret = IsSorted(list, cmp)
			}
		},
	}, {
		Name: "IsSortedStrings",
		Params: []internal.Param{
			{Kind: adt.ListKind},
		},
		Result: adt.BoolKind,
		Func: func(c *internal.CallCtxt) {
			a := c.StringList(0)
			if c.Do() {
				c.Ret = IsSortedStrings(a)
			}
		},
	}},
	CUE: `{
	Comparer: {
		T:    _
		x:    T
		y:    T
		less: bool
	}
	Ascending: {
		Comparer
		T:    number | string
		x:    T
		y:    T
		less: true && x < y
	}
	Descending: {
		Comparer
		T:    number | string
		x:    T
		y:    T
		less: x > y
	}
}`,
}
