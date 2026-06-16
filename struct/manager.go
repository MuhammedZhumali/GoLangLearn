package main

import "fmt"

type TaskManager struct {
	tasks  []Task
	nextID int
}

func NewTaskManager() TaskManager {
	tasks := []Task{}
	return TaskManager{tasks: tasks, nextID: 1}
}

func (tm *TaskManager) AddTask(title string) {
	newTask := Task{ID: tm.nextID, Title: title, Done: false}
	tm.tasks = append(tm.tasks, newTask)
	tm.nextID++
}

func (tm *TaskManager) MarkDone(id int) error {
	for i := range tm.tasks {
		if tm.tasks[i].ID == id {
			tm.tasks[i].Done = true
			return nil
		}
	}

	return fmt.Errorf("Task with ID %d was not found", id)
}

func (tm *TaskManager) DeleteTask(id int) error {
	for i := range tm.tasks {
		if tm.tasks[i].ID == id {
			tm.tasks = append(tm.tasks[:i], tm.tasks[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("task with ID %d not found", id)
}

func (tm *TaskManager) FindTask(id int) (Task, error) {
	for i := range tm.tasks {
		if tm.tasks[i].ID == id {
			return tm.tasks[i], nil
		}
	}

	return Task{}, fmt.Errorf("task with ID %d not found", id)
}

func (tm *TaskManager) ListDone() []Task {
	doneTasks := []Task{}

	for i := range tm.tasks {
		if tm.tasks[i].Done {
			doneTasks = append(doneTasks, tm.tasks[i])
		}
	}

	return doneTasks
}

func (tm *TaskManager) ListPending() []Task {
	pendingTasks := []Task{}

	for i := range tm.tasks {
		if !tm.tasks[i].Done {
			pendingTasks = append(pendingTasks, tm.tasks[i])
		}
	}

	return pendingTasks
}

func (tm *TaskManager) PrintTasks() {
	for i := range tm.tasks {
		tm.tasks[i].Print()
	}
}
