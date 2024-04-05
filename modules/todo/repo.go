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

	todos = append(todos, Todo{
		Task: "halo",
	})

	return todos, nil
}
