package main

import "fmt"

func main() {
	tm := NewTaskManager()
	tm.AddTask("Learn Go")
	tm.AddTask("Build API")
	tm.AddTask("Write tests")
MenuLoop:
	for {
		fmt.Println("Enter your option: ")
		fmt.Println(" 1. Add \n 2. Mark done \n 3. Delete \n 4. List \n 5. Quit")
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
			if err != nil {
				fmt.Println("Invalid input, please enter a number")
				continue
			}
			err = tm.MarkDone(id)
			if err == nil {
				fmt.Println("Task marked as done.")
			} else {
				fmt.Println("Task not found.")
			}

		case 3:
			fmt.Println("Enter task ID to delete: ")
			var id int
			_, err := fmt.Scanln(&id)
			if err != nil {
				fmt.Println("Invalid input, please enter a number")
				continue
			}
			if err := tm.DeleteTask(id); err == nil {
				fmt.Println("Task deleted.")
			} else {
				fmt.Println("Error deleting task:", err)
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

	items := []Printable{
		Task{ID: 1, Title: "Learn Go"},
		Note{Text: "Check many structs in one interface"},
	}

	PrintAll(items)

	fmt.Println(tm.Count())
	tm.Clear()
	fmt.Println(tm.Count())

	items2 := []Describable{
		Task{ID: 1, Title: "Learn Go", Done: false},
		Note{Text: "Interfaces are behavior"},
	}

	PrintDescriptions(items2)
}
