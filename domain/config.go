package domain

import (
  "log"
  "database/sql"
  "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func ConnectToDb(dbUser string, dbPass string) {
  cfg := mysql.Config{
    // Capture connection properties.
    User: dbUser,
    Passwd: dbPass,
    Net:    "tcp",
    Addr:   "127.0.0.1:3306",
    DBName: "animals",
  }
  // Get a database handle.
  var err error
  Db, err = sql.Open("mysql", cfg.FormatDSN())
  if err != nil {
    log.Println(err)
  }

  pingErr := Db.Ping()
  if pingErr != nil {
    log.Println(pingErr)
  }
  log.Println("Connected to DB")
}

