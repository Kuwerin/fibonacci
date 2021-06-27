package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Kuwerin/fibonacci/internal/model"
	"github.com/Kuwerin/fibonacci/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler{
	return &Handler{
		services: services,
	}
}

func (h *Handler) CalculateFibonacci(w http.ResponseWriter, r *http.Request)  {
	req := &model.Fibonacci{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil{
		h.newError(w, r, http.StatusBadRequest, err)	
		return
	}
}

func (h *Handler) newError(w http.ResponseWriter, r *http.Request, code int, err error) {
	h.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (h *Handler) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}){
	w.WriteHeader(code)
	if data != nil{
		json.NewEncoder(w).Encode(data)
	}
}