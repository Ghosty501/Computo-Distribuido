package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"upcrm.com/interesados/internal/controller/interesados"
	"upcrm.com/interesados/pkg/model"
)

type Handler struct {
	ctrl *interesados.Controller
}

func New(ctrl *interesados.Controller) *Handler {
	return &Handler{ctrl}
}

func (h *Handler) Handler(w http.ResponseWriter, req *http.Request) {
	interesadoID := model.InteresadoID(req.FormValue("id"))

	if interesadoID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	interesadoType := model.InteresadoType(req.FormValue("type"))
	if interesadoType == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch req.Method {
	case http.MethodGet:
		v, err := h.ctrl.GetAggregatedInteresados(req.Context(), interesadoID, interesadoType)
		if err != nil && errors.Is(err, interesados.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if err := json.NewEncoder(w).Encode(v); err != nil {
			log.Printf("Response encode error: %v \n", err)
		}

	case http.MethodPut:
		userID := model.UserID(req.FormValue("userID"))
		v, err := strconv.ParseFloat(req.FormValue("value"), 64)
		if err != nil && errors.Is(err, interesados.ErrNotFound) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := h.ctrl.PutRating(req.Context(), interesadoID, interesadoType, &model.Interesados{UserID: userID, Value: model.InteresadosValue(v)}); err != nil {
			log.Printf("Repository put error: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}
