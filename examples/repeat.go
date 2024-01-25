package main

import (
	"fmt"
	"github.com/hovertank3d/stdfp"
	"github.com/hovertank3d/stdfp/lists/list"
)

func main() {
	y := list.FromSlice([]byte("AMOGUS "))
	yy := stdfp.Repeat(y)

	fmt.Println(string(yy.Slice(0, 99)))
}
