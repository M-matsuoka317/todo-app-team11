package main

import (
	"fmt"
	"encoding/json"
	"os"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

const dataFile = "todo.json"

func loadTasks() ([]Task, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var tasks []Task
	if err := json.NewDecoder(file).Decode(&tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func saveTasks(tasks []Task) error {
	file, err := os.Create(dataFile)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(tasks)
}

func nextID(tasks []Task) int {
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}

func AddTask(title string) {
	var tasks []Task
	tasks, _ = loadTasks()
	var new_id = nextID(tasks)
	var task Task
	task.Title = title
	task.ID = new_id
	task.Done = false
	tasks = append(tasks, task)

	// tasks[new_id-1] = task
	// tasks := []Task{task}
	// // var tasks []Task
	saveTasks(tasks)
	// panic("unimplemented")
}

func ListTasks() {
	var tasks []Task
	tasks, _ = loadTasks()
	for _, v := range tasks{
		if v.Done{
			fmt.Println(v.ID, v.Title, "[x]")
		} else {
			fmt.Println(v.ID, v.Title, "[]")
		}
	}
	// panic("unimplemented")
}

func CompleteTask(id int) {
	panic("unimplemented")
}

func DeleteTask(id int) {
	panic("unimplemented")
}
