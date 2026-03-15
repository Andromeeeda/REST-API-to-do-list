package todo

import "time"

type Task struct {
	Title       string
	Description string
	Completed   bool

	TimeStart     time.Time
	TimeCompleted *time.Time
}

func NewTask(title string, description string) Task {
	return Task{
		Title:         title,
		Description:   description,
		Completed:     false,
		TimeStart:     time.Now(),
		TimeCompleted: nil,
	}
}

func (t *Task) Complete() {
	t.Completed = true

	timeCompleted := time.Now()
	t.TimeCompleted = &timeCompleted
}

func (t *Task) UnComplete() {
	t.Completed = false
	t.TimeCompleted = nil
}

