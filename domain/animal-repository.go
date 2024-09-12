package domain

import (
  "fmt"
  "os"
  "database/sql"
  "go-by-example/model"
  "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func ConnectToDb() {
  dbUser := os.Getenv("DBUSER")
  dbPass := os.Getenv("DBPASS")
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
  db, err = sql.Open("mysql", cfg.FormatDSN())
  if err != nil {
    fmt.Println(err)
  }

  pingErr := db.Ping()
  if pingErr != nil {
    fmt.Println(pingErr)
  }
  fmt.Println("Connected to DB")
}

func getAllAnimals() []model.Animal {
  var animals []model.Animal
  rows, err := db.Query("SELECT * FROM animal")
  if err != nil {
    fmt.Println(err)
  }
  defer rows.Close()
  for rows.Next() {
    var animal model.Animal
    if err := rows.Scan(&animal.Id, &animal.Name, &animal.Age); err != nil {
      fmt.Println(err)
    }
    animals = append(animals, animal)
  }
  return animals
}

func getAnimalById(id string) (model.Animal, error) {
  var animal model.Animal

  row := db.QueryRow("SELECT * FROM animal WHERE id = ?", id)
  if err := row.Scan(&animal.Id, &animal.Name, &animal.Age); err != nil {
    if err == sql.ErrNoRows {
      return animal, fmt.Errorf("Animal with ID %d not fuund", id)
    } else {
      fmt.Println(err)
    }
  }
  return animal, nil
}

