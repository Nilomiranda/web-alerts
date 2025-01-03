package routes

import (
	"encoding/json"
	"net/http"
)


type HealthyResponse struct {
  Online bool
}

func GetHealth(w http.ResponseWriter, r *http.Request) {
  var response = HealthyResponse {
    Online: true,
  }
  
  w.Header().Set("Content-Type", "application/json")

  json.NewEncoder(w).Encode(response)
}
