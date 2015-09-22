package main

import "fmt"

type Talk struct {
	Name        string
	Description string
	Properties  map[string]interface{}
}

func (t *Talk) AddProperty(name string, value interface{}) *Talk {
	t.Properties[name] = value

	return t
}

func NewTalk(name, description string) *Talk {
	return &Talk{
		Name: name,
		Description: description,
		Properties: make(map[string]interface{}),
	}
}

func main() {
	talk := NewTalk("Golang", "A quick tour of golang").
		AddProperty("foo", "bar").
		AddProperty("num", 1)

	fmt.Printf("talk: %+v \n", talk)
}