package main

import (
	"fmt"
	"time"
)

func main() {

	max := 2
	c := make(chan string, max)

	for _, v := range []int{1, 2, 3, 4, 5} {
		go func(v int) {
			c <- fmt.Sprintf("Message #%d", name, v)
		}(v)
	}

}