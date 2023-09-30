package main

// Import used packages, bufio and os are used for the line reader on `reader.ReadString()` to allow for the user to input a string with spaces.
// Base `ReadString()` doesn't read in space characters properly so the task details wouldn't be properly defined.
import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// main function is the entry point of the program. It displays a menu of options to the user and performs the corresponding action based on the user's input.
func main() {
	fmt.Printf("Welcome to your to-do list:\n")
	// Declare menuItem variable to store user input.
	var menuItem int
	// Create a new reader (using the bufio package) to read user input from the console.
	reader := bufio.NewReader(os.Stdin)
	// Begin loop for while app runs (until the user chooses to exit).
	for {
		openMenu()
		// Read user input.
		fmt.Scanln(&menuItem)
		// Perform action based on user input
		switch menuItem {
		case 1:
			// Create a new task
			fmt.Println("Enter task name:")
			name, _ := reader.ReadString('\n')
			fmt.Println("Enter task date:")
			date, _ := reader.ReadString('\n')
			fmt.Println("Enter task status:")
			status, _ := reader.ReadString('\n')
			createNewTask(name, date, status)
		case 2:
			// View all tasks
			viewTasks()
		case 3:
			// Update a task
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
			// Remove a task
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
			// Exit the program
			fmt.Println("Exiting...")
			return
		default:
			// Invalid input
			fmt.Println("Invalid choice. Please choose a number between 1 and 5.")
		}
	}
}

func openMenu() {
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
