package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin) //Initialize the reader function
	const filename = "tasks.json"
	err := LoadTasks(filename)
	if err != nil {
		fmt.Println("Error when loading task:", err)
	}

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Create a task")
		fmt.Println("2. Display tasks")
		fmt.Println("3. Flag a task as over")
		fmt.Println("4. Delete a task")
		fmt.Println("5. Save and leave")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var title, description string
			fmt.Print("Title: ")
			title, _ = reader.ReadString('\n') //Assign the value, and ignore the second one, which is error
    		title = title[:len(title)-1]
			fmt.Print("Desc: ")
			description, _ = reader.ReadString('\n')
   			description = description[:len(description)-1]
			CreateTask(title, description)
		case 2:
			DisplayTasks()
		case 3:
			var index int
			fmt.Print("ID of the task to tag as over: ")
			fmt.Scan(&index)
			MarkTaskAsCompleted(index - 1) //-1 To match 0-based indexing.
		case 4:
			var index int
			fmt.Print("ID of the task to delete: ")
			fmt.Scan(&index)
			DeleteTask(index - 1) //Same.
		case 5:
			err := SaveTasks(filename)
			if err != nil {
				fmt.Println("Error when saving the file:", err)
			}
			fmt.Println("Tasks saved to Tasks.json. Closing.")
			return
		default:
			fmt.Println("Invalid option.")
		}
	}
}
