// 代码生成时间: 2025-09-15 07:26:27
package main

import (
    "fmt"
    "log"
    "time"
)

// Scheduler is the structure that holds the tasks to be scheduled.
type Scheduler struct {
    tasks map[string]*Task
}

// Task is the structure that holds the task to be executed.
type Task struct {
    Name        string
    ExecuteFunc func() error
}

// NewScheduler creates a new instance of the scheduler.
func NewScheduler() *Scheduler {
    return &Scheduler{
        tasks: make(map[string]*Task),
    }
}

// AddTask adds a task to the scheduler.
func (s *Scheduler) AddTask(name string, task *Task) error {
    if _, exists := s.tasks[name]; exists {
        return fmt.Errorf("task with name '%s' already exists", name)
    }
    s.tasks[name] = task
    return nil
}

// RemoveTask removes a task from the scheduler.
func (s *Scheduler) RemoveTask(name string) error {
    if _, exists := s.tasks[name]; !exists {
        return fmt.Errorf("task with name '%s' does not exist", name)
    }
    delete(s.tasks, name)
    return nil
}

// Run starts the scheduler and runs the tasks according to their schedule.
func (s *Scheduler) Run() {
    for {
        for name, task := range s.tasks {
            fmt.Printf("Running task: %s
", name)
            if err := task.ExecuteFunc(); err != nil {
                log.Printf("Error running task '%s': %s
", name, err)
            }
        }
        // Add a sleep duration to simulate the tasks running at intervals.
        // In a real-world scenario, you would use a more sophisticated scheduling mechanism.
        time.Sleep(5 * time.Second)
    }
}

// ExampleTask is an example task that prints a message.
func ExampleTask() error {
    fmt.Println("Executing example task...")
    return nil
}

func main() {
    scheduler := NewScheduler()
    
    // Add an example task to the scheduler.
    exampleTask := &Task{Name: "example", ExecuteFunc: ExampleTask}
    if err := scheduler.AddTask(exampleTask.Name, exampleTask); err != nil {
        log.Fatalf("Failed to add task: %s
", err)
    }
    
    // Run the scheduler in a separate goroutine.
    go scheduler.Run()
    
    // Keep the main goroutine alive to allow the scheduler to run.
    select {}
}
