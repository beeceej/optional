package optional

type optional struct {
	data [1]Any
}

// Any represents an empty interface{}
// friendlier version instead of using interface{} everywhere
type Any interface{}

// Of returns an optional of the value given
func Of(d Any) optional {
	return optional{
		data: [1]Any{d},
	}
}

// OfNil returns an optional container containing nil
func OfNil() optional {
	return optional{
		data: [1]Any{nil},
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
func (o optional) Map(f func(d Any) Any) optional {
	if o.Exist() {
		return optional{
			data: [1]Any{f(o.data[0])},
		}
	}
	return optional{data: [1]Any{nil}}
}

func (o optional) Get() Any {
	return o.data[0]
}
