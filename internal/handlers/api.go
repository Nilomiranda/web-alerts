package handlers

import (
	"fmt"
	"net/http"
	"nilomiranda/web-alerts/internal/routes"
)

type HTTPHandler struct {}
func (httpHandler *HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Handling incoming requests...")
}

func InitiateRouter() {
  http.HandleFunc("/api/health", routes.GetHealth)
}
