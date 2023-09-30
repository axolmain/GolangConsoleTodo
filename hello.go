package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Printf("Welcome to your to-do list:\n")
	var menuItem int
	reader := bufio.NewReader(os.Stdin)
	for {
		menu()
		fmt.Scanln(&menuItem)
		switch menuItem {
		case 1:
			fmt.Println("Enter task name:")
			name, _ := reader.ReadString('\n')
			fmt.Println("Enter task date:")
			date, _ := reader.ReadString('\n')
			fmt.Println("Enter task status:")
			status, _ := reader.ReadString('\n')
			createNewTask(name, date, status)
		case 2:
			viewTasks()
		case 3:
			viewTasks()
			var taskNumber int
			fmt.Println("Enter task number to update:")
			fmt.Scanln(&taskNumber)
			if taskNumber < 1 || taskNumber > len(Tasks) {
				fmt.Println("Invalid task number")
				time.Sleep(2 * time.Second)
				continue
			}
			fmt.Printf("Enter new task name (Current: %s):\n", Tasks[taskNumber-1].name)
			name, _ := reader.ReadString('\n')
			fmt.Printf("Enter new task date (Current: %s):\n", Tasks[taskNumber-1].date)
			date, _ := reader.ReadString('\n')
			fmt.Printf("Enter new task status (Current: %s):\n", Tasks[taskNumber-1].status)
			status, _ := reader.ReadString('\n')
			updateTask(taskNumber-1, name, date, status)
		case 4:
			var taskNumber int
			fmt.Println("Enter task number to remove:")
			fmt.Scanln(&taskNumber)
			if taskNumber < 1 || taskNumber > len(Tasks) {
				fmt.Println("Invalid task number")
				time.Sleep(2 * time.Second)
				continue
			}
			removeTask(taskNumber - 1)
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please choose a number between 1 and 5.")
		}
	}
}

func menu() {
	fmt.Printf("Please choose an action (1-5):\n")
	fmt.Printf("1. Create new task\n")
	fmt.Printf("2. View all tasks\n")
	fmt.Printf("3. Update task\n")
	fmt.Printf("4. Remove task\n")
	fmt.Printf("5. Quit\n")
}

func createNewTask(name, date, status string) {
	Tasks = append(Tasks, Task{name, date, status})
}

func updateTask(taskNumber int, name, date, status string) {
	Tasks[taskNumber] = Task{name, date, status}
}

func removeTask(taskNumber int) {
	Tasks = append(Tasks[:taskNumber], Tasks[taskNumber+1:]...)
}

func viewTasks() {
	if len(Tasks) == 0 {
		fmt.Println("No tasks available.")
		time.Sleep(2 * time.Second)
		return
	}
	for i, task := range Tasks {
		fmt.Printf("%d. Task: %s   Do by: %s   Status: %s\n", i+1, task.name, task.date, task.status)
	}
}

var Tasks []Task

type Task struct {
	name   string
	date   string
	status string
}
