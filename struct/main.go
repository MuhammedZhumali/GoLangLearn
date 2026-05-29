package main

import "fmt"

type Task struct {
	ID    int
	Title string
	Done  bool
}

type TaskManager struct {
	tasks []Task
	nextID int
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

	task, found := findTaskById(tasks, 1)

	if found {
		fmt.Println("Found:", task)
	} else {
		fmt.Println("Task not found")
	}

	for i:= range tasks {
		tasks[i].MarkDone()
		tasks[i].Print()
	}

	tasks[0].Rename("Learn Go in depth")
	tasks[0].Print()

	tm := NewTaskManager()

	tm.AddTask("Learn Go")
	tm.AddTask("Build API")
	tm.AddTask("Write tests")

	for i:= range tm.tasks{
		fmt.Printf("ID: %d, Title: %s, Done: %t\n", tm.tasks[i].ID, tm.tasks[i].Title, tm.tasks[i].Done)
	}
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

func findTaskById(tasks []Task, id int) (Task, bool) {
	for i := range tasks {
		if tasks[i].ID == id {
			return tasks[i], true
		}
	}

	return Task{}, false

}


func (t Task) Print(){
	if t.Done {
		fmt.Printf("[✓] %d: - %s\n", t.ID, t.Title)
	} else {
		fmt.Printf("[ ] %d: - %s\n", t.ID, t.Title)
	}
}

func (t *Task) MarkDone() {
	if !t.Done{
		t.Done = true
	}
}

func (t *Task) Rename(newTitle string){
	t.Title = newTitle
}

func NewTaskManager() TaskManager {
	tasks := []Task{}
	return TaskManager{tasks: tasks, nextID: 1}
}

func (tm *TaskManager) AddTask(title string){
	newTask := Task{ID: tm.nextID, Title: title, Done: false}
	tm.tasks = append(tm.tasks, newTask)
	tm.nextID++
}