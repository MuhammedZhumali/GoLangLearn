package main

import "fmt"

type Task struct {
	ID    int
	Title string
	Done  bool
}

func (t Task) Print() {
	if t.Done {
		fmt.Printf("[x] %d: - %s\n", t.ID, t.Title)
	} else {
		fmt.Printf("[ ] %d: - %s\n", t.ID, t.Title)
	}
}

func (t *Task) MarkDone() {
	if !t.Done {
		t.Done = true
	}
}

func (t *Task) Rename(newTitle string) {
	t.Title = newTitle
}


func (t *Task) Toggle() {
	if t.Done {
		t.Done = false
	} else {
		t.Done = true
	}
}
