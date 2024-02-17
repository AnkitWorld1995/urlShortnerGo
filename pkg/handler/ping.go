package handler

import "net/http"

type PingHandler struct{}

func (h *PingHandler) Greet(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, http.StatusOK, "Ping Go....")
}
