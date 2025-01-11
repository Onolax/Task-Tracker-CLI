package main

import (
	"fmt"
	"os"
)

/*
daTask app has:
1. tasks{id,description,status,createdAt, updatedAt}
2. manages a json file default named tasks.json
3. adds, updates, deletes, lists, change status in this json file
4.
*/
func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "add":
			if len(os.Args) > 2 {
				fmt.Println("Usage: add <description>")
			}
			AddTask(os.Args[2])
			return
		case "list":
			PrintTasks(0)
			return
		case "list-not-done":
			if len(os.Args) < 3 {
				fmt.Println("Usage: list-not-done <task name>")
			}
			PrintTasks(1)
			return
		case "list-in-progress":
			PrintTasks(2)
			return
		case "list-done":
			PrintTasks(3)
			return
		case "delete":
			if len(os.Args) < 3 {
				fmt.Println("Usage: delete <task name>")
			}
			DeleteTask(os.Args[2])
			return
		case "update":
			if len(os.Args) < 4 {
				fmt.Println("Usage: add <task> <description>")
			}
			UpdateTask(os.Args[2], os.Args[3])
			return
		case "mark":
			UpdateTaskStatus(os.Args[2], os.Args[3])
			return
		default:
			fmt.Println("unknown command")
			return
		}
	}

}
