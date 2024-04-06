package todo

import (
	"database/sql"
	"time"
)

type Todo struct {
	Id        int
	Task      string
	Note      sql.NullString
	DoneAt    sql.NullTime `db:"done_at"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
}
