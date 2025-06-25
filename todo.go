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
	fmt.Println(loadTasks())
}

func ListTasks() {
	fmt.Println(loadTasks())
}

func CompleteTask(id int) {
	var tasks []Task
	tasks,_ = loadTasks()
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = true
		}
	}
	saveTasks(tasks)
}

func DeleteTask(id int) {
	var tasks []Task
	var updated []Task
	tasks,_ = loadTasks()
	for _, task := range tasks {
		if task.ID != id {
			updated = append(updated, task)
		}
	}
	saveTasks(updated)
}
