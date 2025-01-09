package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitiateDatabase() {
  fmt.Printf("Starting database intitialization...")

  db, err := sql.Open("sqlite3", "../../database.db")
  if err != nil {
    log.Fatal("Error intiating database", err)
  }

  defer db.Close()

  sqlTestStatement := `
    select * from bootstrap;
  `

  fmt.Printf("Running sample SQL statement %s", sqlTestStatement)
  resultRows, err := db.Query(sqlTestStatement)
  if err != nil {
    log.Printf("Error executing statement %s: %e", sqlTestStatement, err)
    return
  }

  defer resultRows.Close()
  for resultRows.Next() {
    var version string
    err = resultRows.Scan(&version)
    if err != nil {
      log.Fatal("Error getting results from row", err)
    }

    fmt.Printf("Current database version %s", version)
  }
}
