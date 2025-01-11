package main

import (
	"fmt"
	"strconv"
	"time"
)

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func AddTask(description string) error {
	err := HandleFile()
	if err != nil {
		return err
	}
	tasks, err := ExtractTasks()
	if err != nil {
		return err
	}
	id := GiveId(tasks)
	newTask := Task{id, description, "Not-Done", time.Now(), time.Now()}
	tasks = append(tasks, newTask)
	err = WriteTask(tasks)
	if err != nil {
		return err
	}
	return nil
}

func PrintTasks(flag int) error {
	tasks, err := ExtractTasks()
	if err != nil {
		return err
	}
	for _, task := range tasks {
		if flag == 3 && task.Status == "Done" {
			fmt.Println(task.Id, ":", task.Description, ":", task.Status)
		} else if flag == 2 && task.Status == "In-Progress" {
			fmt.Println(task.Id, ":", task.Description, ":", task.Status)
		} else if flag == 1 && task.Status == "Not-Done" {
			fmt.Println(task.Id, ":", task.Description, ":", task.Status)
		} else if flag == 0 {
			fmt.Println(task.Id, ":", task.Description, ":", task.Status)
		}
	}
	return nil
}

func DeleteTask(id string) error {
	i, _ := strconv.Atoi(id)
	if err := HandleFile(); err != nil {
		return err
	}

	tasks, err := ExtractTasks()
	if err != nil {
		return err
	}

	newTasks := make([]Task, 0, len(tasks)-1)
	for _, task := range tasks {
		if task.Id != i {
			newTasks = append(newTasks, task)
		}
	}
	err = WriteTask(newTasks)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTask(id, description string) error {
	if description == "" {
		fmt.Println("description is required")
		return nil
	}
	i, _ := strconv.Atoi(id)
	if err := HandleFile(); err != nil {
		return err
	}
	tasks, err := ExtractTasks()
	if err != nil {
		return err
	}
	for idx := range tasks {
		if tasks[idx].Id == i {
			tasks[idx].Description = description
		}
	}
	err = WriteTask(tasks)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTaskStatus(id, status string) error {
	if status != "In-Progress" && status != "Done" {
		fmt.Println("Status wrong")
		return nil
	}
	i, _ := strconv.Atoi(id)
	if err := HandleFile(); err != nil {
		return err
	}
	tasks, err := ExtractTasks()
	if err != nil {
		return err
	}
	for idx := range tasks {
		if tasks[idx].Id == i {
			tasks[idx].Status = status
		}
	}
	err = WriteTask(tasks)
	if err != nil {
		return err
	}

	return nil
}
