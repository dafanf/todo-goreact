package model

// Task represents a single task in the To-Do list.
type Task struct {
	ID          int
	Title       string
	Description string
	Status      string
}
