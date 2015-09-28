package main

import "fmt"

func main() {

	array := [...]string{"hello", "world"}
	//	array = append(array, "!") // cannot compile

	slice := []string{"hello", "world"}
	slice = append(slice, "!")

	fmt.Printf("array is: %T , slice is %T\n", array, slice)

	fmt.Printf("i: %v \n", slice)
	fmt.Printf("i: %v \n", slice[1:])

	for _, v := range slice {
		fmt.Printf("%s, ", v)
	}

	for _, v := range array {
		fmt.Printf("%s, ", v)
	}
}
