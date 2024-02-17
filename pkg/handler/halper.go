package handler

import (
	"encoding/json"
	"net/http"
)

func writeResponse(w http.ResponseWriter, code int, response any) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Method", "Content-Type")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		panic(err)
	}
}
