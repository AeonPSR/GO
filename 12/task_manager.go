package main

import "fmt"

//Creation, Display, Completion, Deletion.

var tasks []Task

func CreateTask(title, description string) {
	task := Task{Title: title, Description: description, Status: "TO DO"}
	tasks = append(tasks, task)
	fmt.Println("Task succesfuly created.")
}

func DisplayTasks() {
	if len(tasks) == 0 {
		fmt.Println("No available task.")
		return
	}
	for i, task := range tasks {
		fmt.Printf("%d. [%s] %s: %s\n", i+1, task.Status, task.Title, task.Description)
	}
}

func MarkTaskAsCompleted(index int) {
	if index < 0 || index >= len(tasks) {
		fmt.Println("Invalid ID.")
		return
	}
	tasks[index].Status = "Done"
	fmt.Println("Task flagged as done.")
}

func DeleteTask(index int) {
	if index < 0 || index >= len(tasks) {
		fmt.Println("Invalid ID.")
		return
	}
	tasks = append(tasks[:index], tasks[index+1:]...)
	fmt.Println("Task removed.")
}
