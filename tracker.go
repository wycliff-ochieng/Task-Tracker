package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// commands := os.Args[1]
const filename = "task.json"

func main() {
	fmt.Println("Arguments", os.Args)
	if len(os.Args) < 3 {
		fmt.Println("Please Enter command..")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			return
		}
		description := os.Args[2]
		err := addTask(description, filename)
		if err != nil {
			fmt.Println("Some text", err)
			return
		}
		fmt.Println("Task added successfully")
		//
	case "list":
		//
	case "update":
		//
	case "delete":
		//
	default:
		fmt.Println("Unknown Command....", command)

	}
}

type Task struct {
	id            int32
	title         string
	completed     bool
	creation_time time.Time
}

func loadTask(filename string) ([]Task, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return []Task{}, nil
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var task []Task
	err = json.Unmarshal(data, &task)
	if err != nil {
		return nil, err
	}
	return task, err
}

func saveTask(filename string, task []Task) error {
	file, err := json.MarshalIndent(task, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		return err
	}
	return nil
}

func addTask(description string, filename string) error {
	if description == "" {
		fmt.Errorf("Description cannot be empty")
	}
	task, err := loadTask(filename)
	if err != nil {
		return nil
	}
	fmt.Println("This is a valid desc", description)
	fmt.Println("Task loaded successfuly", len(task))

	return nil
}
