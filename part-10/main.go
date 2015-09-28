package main

import "fmt"

func main() {

	c := make(chan string)

	go func(c chan string) {
		c <- "hello\n"
		c <- "world\n"
	}(c)

	for {
		fmt.Print(<-c)
	}
}