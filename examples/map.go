package main

import (
	"bytes"
	"fmt"
	"github.com/hovertank3d/stdfp"
	"github.com/hovertank3d/stdfp/lists/list"
	"io"
	"log"
	"os"
	"unicode"
)

func main() {

	f, err := os.Open("repeat.go")
	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	y := list.FromSlice(bytes.Runes(data))
	y = stdfp.MapF(y, unicode.SimpleFold)

	fmt.Println(string(y.Slice(0, -1)))
}
