package todo

import (
	"time"
)

type TodoResponse struct {
	Id        int       `json:"id"`
	Task      string    `json:"task"`
	Note      string    `json:"note"`
	DoneAt    time.Time `json:"done_at"`
	CreatedAt time.Time `json:"created_at"`
}
