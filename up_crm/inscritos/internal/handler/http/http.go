package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"upcrm.com/inscritos/internal/controller/inscritos"
	"upcrm.com/inscritos/pkg/model"
)

type Handler struct {
	ctrl *inscritos.Controller
}

func New(ctrl *inscritos.Controller) *Handler {
	return &Handler{ctrl}
}

func (h *Handler) Handler(w http.ResponseWriter, req *http.Request) {
	inscritosID := model.InscritoID(req.FormValue("id"))

	if inscritosID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	inscritosType := model.InscritoType(req.FormValue("type"))
	if inscritosType == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch req.Method {
	case http.MethodGet:
		v, err := h.ctrl.GetAggregatedInscritos(req.Context(), inscritosID, inscritosType)
		if err != nil && errors.Is(err, inscritos.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if err := json.NewEncoder(w).Encode(v); err != nil {
			log.Printf("Response encode error: %v \n", err)
		}

	case http.MethodPut:
		userID := model.UserID(req.FormValue("userID"))
		v, err := strconv.ParseFloat(req.FormValue("value"), 64)
		if err != nil && errors.Is(err, inscritos.ErrNotFound) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := h.ctrl.PutRating(req.Context(), inscritosID, inscritosType, &model.Inscritos{UserID: userID, Value: model.InscritosValue(v)}); err != nil {
			log.Printf("Repository put error: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}
