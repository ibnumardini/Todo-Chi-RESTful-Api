package todo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type handler struct {
	service *service
}

func newHandler(service *service) handler {
	return handler{service}
}

func (h *handler) FindAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	datas, err := h.service.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, map[string]interface{}{
			"message": "error",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": datas,
	})
}
