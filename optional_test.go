package optional

import (
	"testing"
)

func TestOptionalContainingNilPanics(t *testing.T) {

	testMapOnOptionalOfNilExpectPanic := func() {
		o := Of(nil)
		if !doesMapOperationPanic(func(a Any) Any { return a }, o) {
			t.Fail()
		}
	}
	testMapOnOptionalContainingValue := func() {
		o := Of("a")
		if doesMapOperationPanic(func(a Any) Any { return a }, o) {
			t.Fail()
		}
		res := o.Map(func(a Any) Any { return a.(string) + "b" }).Get()
		if res != "ab" {
			t.Fail()
		}
	}

	testMapOnOptionalOfNilExpectPanic()

	testMapOnOptionalContainingValue()
}

func doesMapOperationPanic(mapFunc func(Any) Any, o optional) (panics bool) {
	defer func() {
		if r := recover(); r != nil {
			panics = true
		}
	}()

	o.Map(mapFunc)

	return panics
}
