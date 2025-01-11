package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
)

func HandleFile() error {
	_, err := os.Stat("tasks.json")
	if os.IsNotExist(err) {
		file, err := os.Create("tasks.json")
		if err != nil {
			return err
		}
		defer file.Close()
	}
	return nil
}

func ExtractTasks() ([]Task, error) {
	file, err := os.Open("tasks.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var tasks []Task
	data, _ := io.ReadAll(file)
	if len(data) > 0 {
		err = json.Unmarshal(data, &tasks)
		if err != nil {
			return nil, err
		}
	}
	return tasks, nil
}

func GiveId(tasks []Task) int {
	var id int
	flag := true
	for {
		id = rand.Intn(999-100) + 100
		for _, task := range tasks {
			if task.Id == id {
				flag = false
				break
			}
		}
		if flag {
			break
		}
	}
	return id
}

func WriteTask(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "\t")
	if err != nil {
		return err
	}
	return os.WriteFile("tasks.json", data, 0644)
}

func displayHelp() {
	fmt.Println("Available Commands:")
	fmt.Println("  add <description>          - Add a new task with the given description")
	fmt.Println("  list                       - List all tasks")
	fmt.Println("  list-not-done              - List tasks with status 'Not Done'")
	fmt.Println("  list-in-progress           - List tasks with status 'In Progress'")
	fmt.Println("  list-done                  - List tasks with status 'Done'")
	fmt.Println("  delete <task_id>           - Delete a task by its ID")
	fmt.Println("  update <task_id> <desc>    - Update the description of a task")
	fmt.Println("  mark <task_id> <status>    - Update the status of a task (e.g., 'Not Done', 'In Progress', 'Done')")
	fmt.Println("  help                       - Display this help message")
}
