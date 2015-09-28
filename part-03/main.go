package main

import "fmt"

func main() {

	i := 0
	for {
		if i > 9 {
			fmt.Printf("i: %d, type: %T\n", i, i);
			break
		}

		i++
	}

	for i :=0 ; i < 20; i++ { fmt.Printf("%d ", i) }
	fmt.Printf("\ni: %d \n", i)

	for  ; i < 20; i++ { fmt.Printf("%d ", i) }
	fmt.Printf("\ni: %d \n", i)
}
