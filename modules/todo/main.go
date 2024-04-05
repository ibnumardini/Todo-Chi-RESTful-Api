package todo

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func LoadModule(dbConn *sqlx.DB) *chi.Mux {
	repo := newRepo(dbConn)
	service := newService(&repo)
	handler := newHandler(&service)

	r := newRouter(&handler)

	return r
}
