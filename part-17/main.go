package main

import (
	"fmt"; "encoding/json"
)

type User struct {
	Login string `json:"login"`
}

type Parent struct {
	Name  string `json:"parent_name"`
	Email string `json:"email"`
}

func (p *Parent) GetName() string {
	return p.Name
}

type Child struct {
	User
	Parent             `json:"parent"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c *Child) GetName() string {
	return c.Name
}

func (c *Child) GetParentName() string {
	return c.Parent.Name
}

func main() {
	child := new(Child)
	child.Login = "child"
	child.Name = "Child Name"
	child.Parent.Name = "Parent Name"
	child.Description = "Child Description"
	child.Email = "child@email.com"

	data, _ := json.Marshal(child)

	fmt.Printf("Marshal: %s\n", data)
	fmt.Printf("Child Name: %s\n", child.GetName())
	fmt.Printf("Parent Name: %s\n", child.GetParentName())
}