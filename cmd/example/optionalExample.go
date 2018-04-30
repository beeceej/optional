package main

import (
	"fmt"

	opt "github.com/beeceej/optional"
)

func appendBytes(v opt.Any) opt.Any {
	return append(v.([]byte), 0234)
}

func main() {
	exist :=
		opt.
			Of([]byte{1}).
			Map(appendBytes).
			Get()

	fmt.Println(exist)
}
