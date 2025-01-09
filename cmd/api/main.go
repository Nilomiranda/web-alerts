package main

import (
	"fmt"
	"net/http"
	"nilomiranda/web-alerts/internal/config"
	"nilomiranda/web-alerts/internal/handlers"

	log "github.com/sirupsen/logrus"
)

func main() {
  log.SetReportCaller(true)
  fmt.Println("Starting server.")

  config.InitiateDatabase()

  handlers.InitiateRouter()

  err := http.ListenAndServe(":8000", nil)

  if err != nil {
    log.Error(err)
  }
}
