package main

import (
	"fmt"

	arrayslicemaps "github.com/Shakhrik/problem-solving/array_slice_maps"
)

func main() {
	a := []int32{0, 1, 2, 4, 6, 5, 3}
	b := arrayslicemaps.FindMedian(a)
	fmt.Println(b)
}
