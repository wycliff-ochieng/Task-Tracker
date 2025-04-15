package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	fmt.Println("Starting in a few...... ")
}

type Task struct{
	id int32
	title string
	completed bool
	creation_time time.Time
}

func loadTask(filename string)([]Task,error){
	if _,err := os.Stat(filename);os.IsNotExist(err){
		return []Task{},nil
	}
	data,err := os.ReadFile(filename)
	if err != nil{
		return nil,err
	}
	var task []Task
	err = json.Unmarshal(data,&task)
	if err != nil{
		return nil,err
	}
	return task,err
}

func addTask(description string, filename string) error{
	if description == ""{
		fmt.Errorf("Description cannot be empty")
	}
	task,err := loadTask(filename)
	if err != nil{
		return  nil
	}
	fmt.Println("This is a valid desc", description)
	fmt.Println("Task loaded successfuly",len(task))

	return nil
}
