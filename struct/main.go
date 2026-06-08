package main

import "fmt"

type Task struct {
	ID    int
	Title string
	Done  bool
}

type TaskManager struct {
	tasks  []Task
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

	task, found := findTaskByID(tasks, 1)
	if found {
		fmt.Println("Found:", task)
	} else {
		fmt.Println("Task not found")
	}

	for i := range tasks {
		tasks[i].MarkDone()
		tasks[i].Print()
	}

	tasks[0].Rename("Learn Go in depth")
	tasks[0].Print()

	tm := NewTaskManager()

	tm.AddTask("Learn Go")
	tm.AddTask("Build API")
	tm.AddTask("Write tests")

	for i := range tm.tasks {
		fmt.Printf("ID: %d, Title: %s, Done: %t\n", tm.tasks[i].ID, tm.tasks[i].Title, tm.tasks[i].Done)
	}

	tm.MarkDone(2)
	fmt.Println("Done:", tm.ListDone())
	fmt.Println("Pending:", tm.ListPending())

	task, found = tm.FindTask(1)
	fmt.Println("Find task 1:", task, found)

	fmt.Println("Deleted task 1:", tm.DeleteTask(1))
	fmt.Println("After delete:", tm.tasks)

MenuLoop:
	for {
		fmt.Println("Enter your option: ")
		fmt.Println("1. Add \n 2. Mark done \n 3. Delete \n 4. List \n 5. Quit")
		var option int
		_, err := fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Invalid input, please enter a number")
			continue
		}

		switch option {
		case 1:
			fmt.Println("Enter task title: ")
			var title string
			fmt.Scanln(&title)
			tm.AddTask(title)

		case 2:
			fmt.Println("Enter task ID to mark done: ")
			var id int
			_, err := fmt.Scanln(&id)
			if err != nil{
				fmt.Println("Invalid input, please enter a number")
				continue
			}
			if tm.MarkDone(id) {
				fmt.Println("Task marked as done.")
			} else {
				fmt.Println("Task not found.")
			}

		case 3:
			fmt.Println("Enter task ID to delete: ")
			var id int
			_, err := fmt.Scanln(&id)
			if err != nil{
				fmt.Println("Invalid input, please enter a number")
				continue
			}
			if tm.DeleteTask(id) {
				fmt.Println("Task deleted.")
			} else {
				fmt.Println("Task not found.")
			}
		case 4:
			fmt.Println("Tasks:: ")
			tm.PrintTasks()
		case 5:
			fmt.Println("Exiting...")
			break MenuLoop
		default:
			fmt.Println("Invalid option")
		}
	}
}

func addTask(tasks []Task, title string) []Task {
	newID := len(tasks) + 1
	newTask := Task{ID: newID, Title: title, Done: false}

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
		if tasks[i].Done {
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

func findTaskByID(tasks []Task, id int) (Task, bool) {
	for i := range tasks {
		if tasks[i].ID == id {
			return tasks[i], true
		}
	}

	return Task{}, false
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

func NewTaskManager() TaskManager {
	tasks := []Task{}
	return TaskManager{tasks: tasks, nextID: 1}
}

func (tm *TaskManager) AddTask(title string) {
	newTask := Task{ID: tm.nextID, Title: title, Done: false}
	tm.tasks = append(tm.tasks, newTask)
	tm.nextID++
}

func (tm *TaskManager) MarkDone(id int) bool {
	for i := range tm.tasks {
		if tm.tasks[i].ID == id {
			tm.tasks[i].Done = true
			return true
		}
	}

	return false
}

func (tm *TaskManager) DeleteTask(id int) bool {
	for i := range tm.tasks {
		if tm.tasks[i].ID == id {
			tm.tasks = append(tm.tasks[:i], tm.tasks[i+1:]...)
			return true
		}
	}

	return false
}

func (tm *TaskManager) FindTask(id int) (Task, bool) {
	for i := range tm.tasks {
		if tm.tasks[i].ID == id {
			return tm.tasks[i], true
		}
	}

	return Task{}, false
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
