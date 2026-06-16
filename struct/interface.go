package main

import "fmt"

type Printable interface {
	Print()
}

type Note struct {
	Text string
}

func (n Note) Print() {
	fmt.Println(n.Text)
}

func PrintAll(items []Printable) {
	for _, item := range items {
		item.Print()
	}
}
