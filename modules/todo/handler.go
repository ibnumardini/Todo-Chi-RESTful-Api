package todo

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

type handler struct {
	service *service
}

func newHandler(service *service) handler {
	return handler{service}
}

func (h *handler) FindAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	todos, err := h.service.FindAll()
	if err != nil {
		log.Error().Err(err).Msg("failed to get all todos")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": todos,
	})
}
