package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"log"
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

	for ans := ""; ans != "N" && ans != "n"; {
		fmt.Printf("Select an option: ")
		fmt.Scanln(&optionSelected)

		switch optionSelected {
		case 1:
			taskNum++
			
			fmt.Print("Add new task: ")
			reader := bufio.NewReader(os.Stdin)
			taskName, _ := reader.ReadString('\n')

			database[taskNum] = Task{
				Name: taskName,
				Status: "Ongoing",
			}
	
			fmt.Println("Task Added!")
			fmt.Println()
		case 2:
			for i := 0; i < len(database); i++ {
				fmt.Printf("%v. %+v\n", i+1, database[i+1].Name)
			}
			fmt.Println()
		case 3:
			for i, v := range database {
				fmt.Printf("%v. %+v\n", i, v.Name)
			}
			fmt.Println()
			fmt.Println("Task ID to delete: ")
			reader := bufio.NewReader(os.Stdin)
			id, _ := reader.ReadString('\n')
			ID, err := strconv.Atoi(strings.TrimSpace(id))
			if err != nil {
				log.Fatal("Enter a number")
			}

			_, ok := database[ID]
			if !ok {
				log.Fatal("Incorrect ID")
			}

			delete(database, ID)
			fmt.Println("Deletion Successful")
		}

		fmt.Print("Would you like to perform another action (Y/N): ")
		fmt.Scanln(&ans)
	}
	
}