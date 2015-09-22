package main

import "fmt"


func main() {
	pos := make([]int, 5) // or pos := []int{1,2,3,4}
	pos[0] = 0; pos[1] = 1; pos[2] = 2; pos[3] = 3; pos[4] = 4

	fmt.Printf("pos: %v \n", pos)

	sub := pos[1:3]
	fmt.Printf("sub: %v \n", sub)

	sub[0] = 5
	fmt.Printf("sub: %v \n", sub)
	fmt.Printf("pos: %v \n", pos)

	sub = make([]int, 5)
	copy(sub, pos)
	pos[0] = 5
	fmt.Printf("sub: %v \n", sub)
	fmt.Printf("pos: %v \n", pos)
}
