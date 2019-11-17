package main

import (
	"fmt"

	"github.com/qdm12/reprint"
)

func main() {
	one := 1
	two := 2
	type myType struct{ A *int }

	// reprint.FromTo usage:
	x := &myType{&one}
	y := new(myType)
	reprint.FromTo(x, y)
	y.A = &two
	fmt.Println(x.A, *x.A) // 0xc0000a0010 1
	fmt.Println(y.A, *y.A) // 0xc0000a0018 2

	// reprint.This usage:
	x2 := myType{&one}
	out := reprint.This(x2)
	y2 := out.(myType)
	y2.A = &two
	fmt.Println(x2.A, *x2.A) // 0xc0000a0010 1
	fmt.Println(y2.A, *y2.A) // 0xc0000a0018 2
}
