package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gocolly/colly/v2"
)

type BritishColumbiaNewsResponse struct {
  LastNewsAt string
}

func GetBrisithColumbiaNews(w http.ResponseWriter, r *http.Request) {
  var lastNewsAt string
  collector := colly.NewCollector()

  collector.OnHTML("div.container div.row div.col-lg-9 p strong", func(h *colly.HTMLElement) {
    fmt.Println("Checking for updated news...", h.Text, lastNewsAt)
    elementValue := h.Text


    if lastNewsAt == "" {
      lastNewsAt = elementValue
    }

    if (lastNewsAt != "") {
      w.Header().Set("Content-Type", "application/json")

      response := BritishColumbiaNewsResponse {
        LastNewsAt: lastNewsAt, 
      }

      json.NewEncoder(w).Encode(response)

      return
    }

  })

  collector.OnRequest(func(r *colly.Request) {
    fmt.Println("Accessing URL:", r.URL)
  })

  collector.Visit("https://www.welcomebc.ca/immigrate-to-b-c/news")
}
