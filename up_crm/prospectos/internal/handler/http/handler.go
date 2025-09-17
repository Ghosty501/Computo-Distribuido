package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"upcrm.com/prospectos/internal/controller/prospectos"
	"upcrm.com/prospectos/internal/repository"
)

type Handler struct {
	ctrl *prospectos.Controller
}

func New(ctrl *prospectos.Controller) *Handler {
	return &Handler{ctrl}
}

func (h *Handler) GetProspecto(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	ctx := req.Context()
	m, err := h.ctrl.Get(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Repository error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(m); err != nil {
		log.Printf("Response error: v%\n", err)
	}
}
