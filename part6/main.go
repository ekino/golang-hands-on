package main

import "fmt"

type Talk struct {
	Name        string
	Description string
	Properties  map[string]interface{}
}

func main() {
	talk := Talk{
		Name: "Golang",
		Description: "A quick tour of golang",
		Properties: map[string]interface{}{
			"foo": "bar",
			"hello": "world!",
			"num": 1,
		},
	}

	fmt.Printf("talk: %+v \n", talk)
}
