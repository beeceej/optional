# Implementation of Optionals in golang...

Testing out the idea of utilizing a fixed size array as a container denoting the optional presence of a value.

uses `interface{}` which from what I can tell is not used in other libraries. Other libraries seem to utilize code generation
to cope for the lack of typesafe generics.

This library will fail at runtime if the `Any` type is dereferenced with an incorrect type assertion.
Use in production at your own peril. This is more an excercise in determining how useful an Optional type could be in go.

## Todo...

* Tests?
* Ideas? Please contribute!