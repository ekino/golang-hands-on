package main

import "fmt"

func main() {
	r := [...]int{1,2,3,4,5}

	for _, v := range r {
		go func() {
			fmt.Printf("Error: Square of %d is %d\n", v, v*v)
		}()
	}


	for _, v := range r {
		go func(v int) {
			fmt.Printf("Good: Square of %d is %d\n", v, v*v)
		}(v)
	}

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}