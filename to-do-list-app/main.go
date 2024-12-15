package main

import (
	"fmt"
)

type Task struct {
	Name string
	Status string
}
func main() {
	fmt.Println("------ To-Do ------")
	fmt.Println("1. Add Task")
	fmt.Println("2. View Tasks")
	fmt.Println("3. Delete Tasks")

	database := make(map[int]Task)
	taskNum := 0
	var optionSelected int

	for {
		fmt.Printf("Select an option: ")
		fmt.Scanln(&optionSelected)

		switch optionSelected {
		case 1:
			taskNum++
			var newTask string
			fmt.Print("Add new task: ")
			fmt.Scanln(&newTask)
	
			database[taskNum] = Task{
				Name: newTask,
				Status: "Ongoing",
			}
	
			fmt.Println("Task Added!")
		case 2:
			for i, v := range database {
				fmt.Printf("%v. %v\n", i, v)
			}
		}
	}
	
}