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

type Describable interface {
	Description() string
}

func (n Note) Description() string {
	return fmt.Sprintf("Note: %s", n.Text)
}

func PrintDescriptions(items []Describable) {
	for _, item := range items {
		fmt.Println(item.Description())
	}
}
