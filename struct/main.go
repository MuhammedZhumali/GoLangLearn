package main

import "fmt"

type Task struct {
	ID    int
	Title string
	Done  bool
}

func main() {
	tasks := []Task{
		{ID: 1, Title: "Learn Go", Done: true},
		{ID: 2, Title: "Leetcode solve", Done: true},
		{ID: 3, Title: "Backend api build", Done: false},
	}
	tasks = addTask(tasks, "New task check")

	tasks = markDone(tasks, 3)
	for i := 0; i < len(tasks); i++ {
		fmt.Println(tasks[i])
	}

	doneTasks := filterDone(tasks)
	fmt.Println(doneTasks)

	fmt.Println(tasks)
	tasks = deleteTask(tasks, 2)
	fmt.Println(tasks)
}

func addTask(tasks []Task, title string) []Task {
	newId := len(tasks) + 1

	newTask := Task{newId, title, false}

	tasks = append(tasks, newTask)

	return tasks
}

func markDone(tasks []Task, id int) []Task {

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = true
		}
	}

	return tasks
}

func filterDone(tasks []Task) []Task {
	result := []Task{}

	for i := range tasks {
		if tasks[i].Done == true {
			result = append(result, tasks[i])
		}
	}
	return result
}

func deleteTask(tasks []Task, id int) []Task {
	result := []Task{}

	for i := range tasks {
		if tasks[i].ID != id {
			result = append(result, tasks[i])
		}
	}
	return result
}
