package exprtree

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	want := 5
	expr := New(5)

	if reflect.TypeOf(expr) != reflect.TypeOf(&ExprTree{}) {
		t.Errorf("expr is not of Type: %T", ExprTree{})
	}

	if expr.Value != float64(want) {
		t.Errorf("Value %f is not equal %f", expr.Value, float64(want))
	}
}

func TestInsertLeft(t *testing.T) {

}

func Test(t *testing.T) {
	//(((2) 4 (6)) 1 (5 (3)))
	want := "(((2) 4 (6)) 1 (5 (3)))"
	expr := New(1)

	expr.InsertLeft(2)
	expr.InsertRight(3)
	expr.InsertLeft(4)
	expr.Left.InsertRight(6)
	expr.InsertRight(5)

	tree := expr.Print()
	if tree != want {
		t.Errorf("Value %s is not equal to %s", tree, want)
	}
}
