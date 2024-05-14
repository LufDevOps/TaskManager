package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	// "strings"
)

type Task struct {
    Description string
    Completed   bool
}

type TaskSlice struct {
    tasks []Task
}

func NewTaskSlice() *TaskSlice {
    return &TaskSlice{tasks: []Task{}}
}

func (tm *TaskSlice) AddTask(description string) {
    for _, task := range tm.tasks {
        if task.Description == description {
            fmt.Println("Task description already exists.")
            return
        }
    }
    // If not duplicate, add the task
    tm.tasks = append(tm.tasks, Task{Description: description, Completed: false})
    fmt.Println("Task added successfully.")
}


func (ts *TaskSlice) CompleteTask(i int) {
    if i >= 0 && i < len(ts.tasks) {
        ts.tasks[i].Completed = true
    } else {
        fmt.Println("task not found.")
    }
}

func (ts *TaskSlice) ListTasks() {
    for i, task := range ts.tasks {
        status := "pending"
        if task.Completed {
            status = "completed"
        }
        fmt.Printf("%d. %s [%s]\n", i+1, task.Description, status)
    }
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    TaskSlice := NewTaskSlice()

    for {
    	fmt.Println(`
	Please choose an option:
	1. Add task
	2. Complete task
	3. List tasks
	Your option: `)

        scanner.Scan()
        option := scanner.Text()

        if option == "1" {
            fmt.Print("Enter task description: ")
            scanner.Scan()
            description := scanner.Text()
			// for i := 0; i < len(TaskSlice.tasks); i++ {
			// 	if TaskSlice.tasks[i].Description == strings.ToLower(description) {
			// 		fmt.Println("Duplicate Task Description with Task: ",i+1,". ",TaskSlice.tasks[i].Description)
			// 		continue
			// 	}
			// }
			TaskSlice.AddTask(description)
            
        } else if option == "2" {
			if len(TaskSlice.tasks)==0 {
				fmt.Println("!Task slice is empty, add task first")
				continue
			}
            fmt.Print("Please enter task id to complete: ")
            scanner.Scan()
            i, err := strconv.Atoi(scanner.Text())
            if err != nil {
                fmt.Println("Invalid input for task id.")
                continue
            }
            TaskSlice.CompleteTask(i - 1)
        } else if option == "3" {
            TaskSlice.ListTasks()
        } else {
            fmt.Println("Invalid option, please choose again.")
        }
    }
}