// Copyright (c) 2011 chuckldev.  All rights reserved.

package exprtree

import "fmt"

type ExprTree struct {
	Value float64
	Left  *ExprTree
	Right *ExprTree
}

func New(value float64) *ExprTree {
	tree := &ExprTree{Value: value}
	return tree
}

func (e *ExprTree) InsertLeft(value float64) {
	if e == nil {
		e.Left = &ExprTree{Value: value}
	} else {
		tree := &ExprTree{Value: value}
		tree.Left = e.Left
		e.Left = tree
	}
}

func (e *ExprTree) InsertRight(value float64) {
	if e == nil {
		e.Right = &ExprTree{Value: value}
	} else {
		tree := &ExprTree{Value: value}
		tree.Right = e.Right
		e.Right = tree
	}
}

func (e *ExprTree) Print() string {
	// Copyright (c) 2011 The Go Authors. All rights reserved.
	if e == nil {
		return "||"
	}

	s := ""

	if e.Left != nil {
		s += e.Left.Print() + " "
	}

	s += fmt.Sprint(e.Value)

	if e.Right != nil {
		s += " " + e.Right.Print()
	}

	return "|" + s + "|"
}
