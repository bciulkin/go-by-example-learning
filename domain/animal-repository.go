package domain

import (
  "log"
  "os"
  "database/sql"
  "go-by-example/model"
  "github.com/go-sql-driver/mysql"
  "fmt"
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
    log.Println(err)
  }

  pingErr := db.Ping()
  if pingErr != nil {
    log.Println(pingErr)
  }
  log.Println("Connected to DB")
}

func getAllAnimals() ([]model.Animal, error) {
  var animals []model.Animal
  rows, err := db.Query("SELECT * FROM animal")
  if err != nil {
    return animals, fmt.Errorf("getAllAnimals: %v", err)
  }
  defer rows.Close()
  for rows.Next() {
    var animal model.Animal
    if err := rows.Scan(&animal.Id, &animal.Name, &animal.Age); err != nil {
      return animals, fmt.Errorf("getAllAnimals: %v", err)
    }
    animals = append(animals, animal)
  }
  return animals, nil
}

func getAnimalById(id string) (model.Animal, error) {
  var animal model.Animal

  row := db.QueryRow("SELECT * FROM animal WHERE id = ?", id)
  if err := row.Scan(&animal.Id, &animal.Name, &animal.Age); err != nil {
    if err == sql.ErrNoRows {
      return animal, fmt.Errorf("Animal with ID not found: %s", id)
    } else {
      log.Println(err)
    }
  }
  return animal, nil
}

func addAnimal(anml model.Animal) (model.Animal, error) {
  _, err := db.Exec("INSERT INTO animal (id, name, age) VALUES (?, ?, ?)", anml.Id, anml.Name, anml.Age)
  if err != nil {
    return anml, fmt.Errorf("addAnimal: %v", err)
  }
  return getAnimalById(anml.Id.String())
  
}

func updateAnimal(anml model.Animal) (model.Animal, error) {
  _, err := db.Exec("UPDATE animal SET name = ?, age = ? WHERE id = ?", anml.Name, anml.Age, anml.Id)
  if err != nil {
    return anml, fmt.Errorf("updateAnimal: %v", err)
  }
  return getAnimalById(anml.Id.String())
}

func deleteAnimal(id string) (string, error) {
  _, err := db.Exec("DELETE FROM animal WHERE id = ?", id)
  if err != nil {
    return id, fmt.Errorf("deleteAnimal: %v", err)
  }
  return id, nil
}
