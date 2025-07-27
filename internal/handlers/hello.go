package handlers

import (
	"encoding/json"
	"net/http"

	"any-api/internal/services"
)

type HelloHandler struct {
	Service *services.HelloService
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	message := h.Service.GetHelloMessage()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}