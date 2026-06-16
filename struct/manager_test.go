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

func TestDeleteTask(t *testing.T) {
	tm := NewTaskManager()
	tm.AddTask("Task test 1")
	tm.AddTask("Task test 2")

	err := tm.DeleteTask(1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(tm.tasks) != 1 {
		t.Fatalf("Expected 1 task after deletion, got %d", len(tm.tasks))
	}
	if tm.tasks[0].ID != 2 {
		t.Errorf("Expected remaining task ID 2, got %d", tm.tasks[0].ID)
	}
	err = tm.DeleteTask(999)
	if err == nil {
		t.Errorf("Expected error when deleting non-existent task, got nil")
	}
}

func TestMarkDone(t *testing.T) {
	tm := NewTaskManager()
	tm.AddTask("Task test 1")
	err := tm.MarkDone(1)
	if err != nil {
		t.Errorf("Expected no error when marking task as done, got %v", err)
	}
	if tm.tasks[0].Done != true {
		t.Errorf("Expected task done status true, got false")
	}
	err = tm.MarkDone(999)
	if err == nil {
		t.Errorf("Expected error when marking non-existent task as done, got nil")
	}
}

func TestFindTask(t *testing.T) {
	tm := NewTaskManager()
	tm.AddTask("Task test 1")
	_, err := tm.FindTask(1)
	if err != nil {
		t.Errorf("Expected to find task with ID 1, got error %v", err)
	}
	_, err = tm.FindTask(999)
	if err == nil {
		t.Errorf("Expected error when finding non-existent task, got nil")
	}
}

func TestListDoneAndPending(t *testing.T) {
	tm := NewTaskManager()
	tm.AddTask("Task test 1")
	tm.AddTask("Task test 2")
	tm.AddTask("Task test 3")
	tm.MarkDone(2)
	taskDone := tm.ListDone()
	if len(taskDone) != 1 {
		t.Errorf("Expected 1 done task, got %d", len(taskDone))
	}
	if taskDone[0].ID != 2 {
		t.Errorf("Expected done task ID 2, got %d", taskDone[0].ID)
	}
	taskPending := tm.ListPending()
	if len(taskPending) != 2 {
		t.Errorf("Expected 2 pending tasks, got %d", len(taskPending))
	}
}
