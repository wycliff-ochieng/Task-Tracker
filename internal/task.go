package internal

import (
	"fmt"
	"time"
)

type status string

const (
	STATUS_TODO        status = "todo"
	STATUS_IN_PROGRESS status = "in-progress"
	STATUS_DONE        status = "done"
)

type Task struct {
	ID          int32     `json:"id"`
	Description string    `json:"description"`
	Status      status    `json:"status"`
	CreatedAt   time.Time `json:"createdat"`
	UpdatedAt   time.Time `json:"updatedat"`
}

func NewTask(id int32, description string) *Task {
	return &Task{
		ID:          id,
		Description: description,
		Status:      STATUS_TODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func ListTask(status status) error {
	tasks, err := ReadTaskFromList()
	if err != nil {
		return err
	}
	if len(tasks) == 0 {
		fmt.Println("No Tasks Found!!!!")
		return nil
	}

	var filteredTasks []Task
	switch status {
	case "all":
		filteredTasks == tasks
	case STATUS_TODO:
		for _, task := range tasks {
			if task.Status == STATUS_TODO {
				filteredTasks = append(filteredTasks, task)
			}
		}
	case STATUS_IN_PROGRESS:
		for _, task := range tasks {
			if task.Status == STATUS_IN_PROGRESS {
				filteredTasks = append(filteredTasks, task)
			}
		}
	case STATUS_DONE:
		for _, task := range tasks {
			if task.Status == STATUS_DONE {
				filteredTasks = append(filteredTasks, task)
			}
		}
	}
	fmt.Println("Tasks %s", tasks)
	return nil
}

func AddTasks(description string) error {
	tasks, err := ReadTaskFromFile()
	if err != nil {
		return err
	}

	var newTaskID int32
	if len(tasks) > 0 {
		lastTask := tasks[len(tasks)-1]
		newTaskID = lastTask.ID + 1
	} else {
		newTaskID = 0
	}

	task := NewTask(newTaskID, description)
	tasks = append(tasks, *task)
	fmt.Println("Task added succesfully %s", task)
	return writeTaskToFile(task)
}

func UpdateTask(id string) error {
	tasks, err := ReadTaskFromFile()
	if err != nil {
		return err
	}

	var taskExists bool = false
	var updatedTask []Task
	for _, task := range tasks {
		if task.ID == id {
			taskExists = true
			switch status {
			case STATUS_DONE:
				task.Status = STATUS_DONE
			case STATUS_IN_PROGRESS:
				task.Status = STATUS_IN_PROGRESS
			case STATUS_TODO:
				task.Status = STATUS_TODO
			}
			task.UpdatedAt = time.Now()
		}

	}
	updatedTask = append(updatedTask, task)
	if !taskExists {
		fmt.Errorf("Input a valid task")
	}
	fmt.Println("updated suceesfully")
	return writeTaskToFile(task)
}

func ReadTaskFromList()
