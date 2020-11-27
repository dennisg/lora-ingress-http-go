package actuator

import (
	"encoding/json"
	"net/http"
)

func Initialize()  {
	http.HandleFunc("/actuator/health", health)
}

func health(w http.ResponseWriter, r *http.Request)  {
	_ = json.NewEncoder(w).Encode(struct {
		Status string `json:"status"`
	}{"UP"})
}