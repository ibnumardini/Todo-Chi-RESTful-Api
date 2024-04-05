package todo

import "github.com/go-chi/chi/v5"

func newRouter(h *handler) *chi.Mux {
	r := chi.NewMux()

	r.Get("/todo", h.FindAll)

	return r
}
