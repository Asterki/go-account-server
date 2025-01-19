package serv

import (
  "database/sql"
  "fmt"

  _ "github.com/mattn/go-sqlite3"

  log "github.com/sirupsen/logrus"
)

func OpenDB() *sql.DB {
  db, err := sql.Open("sqlite3", "./test.db")
  if err != nil {
    log.Fatal(fmt.Sprintf("Error opening the database: %s", err))
  }
  return db
}

func CloseDB(db *sql.DB) {
  err := db.Close()
  if err != nil {
    log.Fatal(fmt.Sprintf("Error closing the database: %s", err))
  }
}

func CreateTables(db *sql.DB) {
  _, err := db.Exec(`CREATE TABLE IF NOT EXISTS accounts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL,
    password TEXT NOT NULL
  )`)
  if err != nil {
    log.Fatal(fmt.Sprintf("Error creating the accounts table: %s", err))
  }
}
