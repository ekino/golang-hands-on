package main

import "fmt"

const w  =  3

var x = 2

func main() {
	var y int64
	z := 1
	z = 3
	x = 3
	// w = 3 # fail

	fmt.Printf("w: %d, type: %T\n", w, w)
	fmt.Printf("x: %d, type: %T\n", x, x)
	fmt.Printf("y: %d, type: %T\n", y, y)
	fmt.Printf("z: %d, type: %T\n", z, z)
}
