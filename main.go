package main

import (
	"fmt"

	opt "github.com/beeceej/optional/optional"
)

func add1(v opt.Any) opt.Any {
	return append(v.([]byte), 0234)
}

func main() {
	exist :=
		opt.
			Of([]byte{1}).
			Map(add1).
			Get()

	fmt.Println(exist)
}
