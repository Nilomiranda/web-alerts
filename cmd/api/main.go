package main

import (
	"fmt"
	"net/http"
	"nilomiranda/web-alerts/internal/handlers"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
  log.SetReportCaller(true)
  var router *chi.Mux = chi.NewRouter()

  handlers.Handler(router)

  fmt.Println("Starting server.")

  err := http.ListenAndServe("localhost:8000", router)

  if err != nil {
    log.Error(err)
  }
}
