package main

//Just the task itself, if need to be expanded.

type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
