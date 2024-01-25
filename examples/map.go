package main

import (
	"fmt"
	"github.com/hovertank3d/stdfp"
)

func main() {
	test := []byte("amogus")

	x := stdfp.NewListSlice(test)
	fmt.Println(string(x.Slice(0, x.Size())))

	x = x.Map(func(b byte) byte {
		return b + 1
	})
	fmt.Println(string(x.Slice(0, x.Size())))
}
