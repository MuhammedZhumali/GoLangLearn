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

	tests := []struct {
		id      int
		wantErr bool
	}{
		{id: 999, wantErr: true},
		{id: 1000, wantErr: true},
		{id: 1, wantErr: false},
	}
	for _, tt := range tests {
		err := tm.MarkDone(tt.id)
		if tt.wantErr && err == nil {
			t.Errorf("Expected error when marking non-existent task with ID %d as done, got nil", tt.id)
		}
		if !tt.wantErr && err != nil {
			t.Errorf("Expected to mark task with ID %d as done, got error %v", tt.id, err)
		}
		if !tt.wantErr && !tm.tasks[0].Done {
			t.Errorf("Expected task with ID %d to be done", tt.id)
		}
	}
}

func TestFindTask(t *testing.T) {
	tm := NewTaskManager()
	tm.AddTask("Task test 1")

	tests := []struct {
		id      int
		wantErr bool
	}{
		{id: 999, wantErr: true},
		{id: 1000, wantErr: true},
		{id: 1, wantErr: false},
	}
	for _, tt := range tests {
		task, err := tm.FindTask(tt.id)
		if tt.wantErr && err == nil {
			t.Errorf("Expected error when finding non-existent task with ID %d, got nil", tt.id)
		}
		if !tt.wantErr && err != nil {
			t.Errorf("Expected to find task with ID %d, got error %v", tt.id, err)
		}
		if !tt.wantErr && task.ID != tt.id {
			t.Errorf("Expected task with ID %d, got task with ID %d", tt.id, task.ID)
		}
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

func TestCount(t *testing.T) {
	tm := NewTaskManager()
	if tm.Count() != 0 {
		t.Errorf("Expected to be new array of tasks")
	}
	tm.AddTask("Task test 1")
	tm.AddTask("Task test 2")
	if tm.Count() != 2 {
		t.Errorf("Expected to be 2 tasks in array")
	}

	tm.DeleteTask(1)
	if tm.Count() != 1 {
		t.Errorf("After deleting task must be only 1")
	}
}

func TestClear(t *testing.T) {
	tm := NewTaskManager()
	tm.AddTask("Task test 1")
	tm.AddTask("Task test 2")
	tm.Clear()
	if tm.Count() != 0 {
		t.Errorf("Must be 0 elements after Clear")
	}
}
