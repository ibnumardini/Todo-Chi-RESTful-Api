package todo

import "github.com/jmoiron/sqlx"

type Repo interface {
	FindAll() ([]Todo, error)
}

type repo struct {
	db *sqlx.DB
}

func newRepo(db *sqlx.DB) repo {
	return repo{db}
}

func (r *repo) FindAll() ([]Todo, error) {
	var todos []Todo

	err := r.db.Select(&todos, "SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}

	return todos, nil
}
