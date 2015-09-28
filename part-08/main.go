package main

import "fmt"

func main() {

	count := func(name string) int {
		return len(name)
	}

	fmt.Printf("Count: %d\n", count("Golang"))

	hello := "Hello"

	join := func(name string) string {
		return fmt.Sprintf("%s %s", hello, name)
	}

	fmt.Printf("Join: %s\n", join("Golang"))
}