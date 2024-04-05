package todo

import "time"

type Todo struct {
	Task      string
	DoneAt    time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
