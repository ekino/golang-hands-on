package main

import (
	"fmt"
)

type Countable interface {
	Len() int
}

type Paragraph struct {
	Name string
}

func (p *Paragraph) Len() int {
	return len(p.Name)
}

type Fragments struct {
	Paragraphs []*Paragraph
}

func (f *Fragments) Len() int {
	return len(f.Paragraphs)
}

func Count(v Countable) int {
	return v.Len()
}

func main() {
	p := &Paragraph{"Salut comment ca va ?"}
	f := &Fragments{[]*Paragraph{p}}

	fmt.Printf("Paragraph Count: %d\n", Count(p))
	fmt.Printf("Fragment Count: %d\n", Count(f))
}