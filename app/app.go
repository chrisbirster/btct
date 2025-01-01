package app

import "fmt"

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Complete    bool   `json:"complete"`
}

type App struct {
	tasks  map[int]Task
	nextID int
}

// NewApp initializes a new app
func NewApp() *App {
	return &App{
		tasks:  make(map[int]Task),
		nextID: 0,
	}
}

// AddTask adds a new task to the queue.
func (a *App) AddTask(description string) Task {
	a.nextID++
	task := Task{
		ID:          a.nextID,
		Description: description,
		Complete:    false,
	}
	a.tasks[a.nextID] = task
	return task
}

// ListTasks returns all tasks (filtered by completion status if requested).
func (a *App) ListTasks(onlyIncomplete bool) []Task {
	var taskList []Task
	for _, task := range a.tasks {
		if onlyIncomplete && task.Complete {
			continue
		}
		taskList = append(taskList, task)
	}
	return taskList
}

// GetTask retrieves a task from the task queue
func (a *App) GetTask(id int) (Task, error) {
	task, exists := a.tasks[id]
	if !exists {
		return Task{}, fmt.Errorf("task with ID %d not found", id)
	}
	return task, nil
}

// MarkTaskComplete marks a task complete
func (a *App) MarkTaskComplete(id int) error {
	task, exists := a.tasks[id]
	if !exists {
		return fmt.Errorf("task with ID %d not found", id)
	}
	task.Complete = true
	a.tasks[id] = task
	return nil
}
