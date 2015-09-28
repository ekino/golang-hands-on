package main

import "fmt"

func main() {

	ubc := make(chan string, 2)
	ubc <- "hello\n"
	ubc <- "world\n"

	fmt.Print(<-ubc)
	fmt.Print(<-ubc)

	bc := make(chan string)
	bc <- "hello\n"

	fmt.Print(<-bc)
}