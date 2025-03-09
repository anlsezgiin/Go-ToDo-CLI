package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const fileName = "tasks.txt"

func main() {

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please enter a command: add, list or delete")
		return
	}

	command := args[1]

	switch command {
	case "add":
		if len(args) < 3 {
			fmt.Println("Please add task!")
			return
		}
		addTask(args[2])

	case "list":
		listTasks()

	case "delete":
		if len(args) < 3 {
			fmt.Println("Enter the number of the task you want to delete!")
			return
		}
		deleteTask(args[2])

	default:
		fmt.Println("Unknown command:", command)
	}
}

// add task
func addTask(task string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(task + "\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("New task added:", task)
}

// list task
func listTasks() {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	tasks := strings.Split(string(data), "\n")
	fmt.Println("Task List:")
	for i, task := range tasks {
		if task != "" {
			fmt.Printf("%d. %s\n", i+1, task)
		}
	}
}

// delete task
func deleteTask(taskNumber string) {
	// read file
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	tasks := strings.Split(string(data), "\n")

	
	index, err := strconv.Atoi(taskNumber)
	if err != nil || index < 1 || index > len(tasks)-1 {
		fmt.Println("Invalid task number!")
		return
	}

	// delete from slice
	tasks = append(tasks[:index-1], tasks[index:]...)

	// updated data to .txt
	err = os.WriteFile(fileName, []byte(strings.Join(tasks, "\n")), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Task deleted:", taskNumber)
}
