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
	testMapOnOptionalContainingValueExpectFunctionCalled := func() {
		o := Of("a")
		if doesMapOperationPanic(func(a Any) Any { return a }, o) {
			t.Fail()
		}
		res := o.Map(func(a Any) Any { return a.(string) + "b" }).Get()
		if res != "ab" {
			t.Fail()
		}
	}
	testGetOrErrExpectOptionalWithData := func() {
		o := Of(nil)

		opt, err := o.GetOrErr()
		if err == nil || opt.Get() != nil {
			t.Fail()
		}
	}
	testGetOrErrExpectErr := func() {
		o := Of("A")

		opt, err := o.GetOrErr()
		if err != nil || opt.Get() == nil {
			t.Fail()
		}
	}

	testMapOnOptionalOfNilExpectPanic()
	testMapOnOptionalContainingValueExpectFunctionCalled()
	testGetOrErrExpectOptionalWithData()
	testGetOrErrExpectErr()
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
