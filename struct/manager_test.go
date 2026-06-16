package main

import "testing"

func TestAddTask(t *testing.T) {
	tm := NewTaskManager()
	tm.AddTask("Learn GO")
	if len(tm.tasks) != 1 {
		t.Fatalf("Expected 1 task, got %d", len(tm.tasks))
	}
	if tm.tasks[0].ID != 1 {
		t.Errorf("Expected task ID 1, got %d", tm.tasks[0].ID)
	}
	if tm.tasks[0].Title != "Learn GO" {
		t.Errorf("Expected task title 'Learn GO', got '%s'", tm.tasks[0].Title)
	}
	if tm.tasks[0].Done != false {
		t.Errorf("Expected task done status false, got true")
	}
}
