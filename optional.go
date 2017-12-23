package optional

import "errors"

type optional struct {
	data     [1]Any
	canBeNil bool
}

// Any represents an empty interface{}
// friendlier version instead of using interface{} everywhere
type Any interface{}

// mapAny is an alias for the function signature used in a map call
type mapAny func(a Any) Any

// OfNillable returns an optional of the value given which may be nil
func OfNillable(a Any) optional {
	return optional{
		data:     [1]Any{a},
		canBeNil: true,
	}
}

// Of returns an optional container, which may not contain a nil value
func Of(a Any) optional {
	return optional{
		data:     [1]Any{a},
		canBeNil: false,
	}
}

// Exist checks if a value exists in the optional
func (o optional) Exist() bool {
	return o.data[0] != nil
}

// Map applys a function to the optional under the condition that data exists within the optional
// If data exists (as defined by optional.Exist(), then apply the function and return an optional with
// the result of f inside the optional
// If data doesn't exist return an empty optional
func (o optional) Map(m mapAny) optional {
	if o.Exist() {
		return optional{
			data: [1]Any{m(o.data[0])},
		}
	}
	o.panicIfNil()
	return optional{data: [1]Any{nil}}
}

// Get returns data if present
func (o optional) Get() Any {
	return o.data[0]
}

func (o optional) panicIfNil() {
	if !o.canBeNil {
		panic(errors.New("Value wrapped in optional must not be nil"))
	}
}

// Get returns either the data if it exists, or an error if no data in container
func (o optional) GetOrErr() (opt optional, err error) {
	if o.Exist() {
		opt = optional{
			data:     [1]Any{o.data},
			canBeNil: o.canBeNil,
		}
		return opt, err
	}
	err = errors.New("No data contained in optional container")
	return opt, err
}
