package domain

import (
  "log"
  "database/sql"
  "github.com/go-sql-driver/mysql"
)


func ConnectToDb(dbUser string, dbPass string) (*sql.DB) {
  cfg := mysql.Config{
    // Capture connection properties.
    User: dbUser,
    Passwd: dbPass,
    Net:    "tcp",
    Addr:   "127.0.0.1:3306",
    DBName: "animals",
  }

  // Get a database handle.
  db, err := sql.Open("mysql", cfg.FormatDSN())
  if err != nil {
    log.Println(err)
  }

  pingErr := db.Ping()
  if pingErr != nil {
    log.Println(pingErr)
  }
  log.Println("Connected to DB")
  return db
}

