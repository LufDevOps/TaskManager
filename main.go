package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
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

func (ts *TaskSlice) AddTask(description string) {
    ts.tasks = append(ts.tasks, Task{Description: description, Completed: false})
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
            TaskSlice.AddTask(description)
        } else if option == "2" {
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
