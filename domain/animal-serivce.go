package domain

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "net/http"
  "os"
  "database/sql"
  "go-by-example/model"
  "github.com/go-sql-driver/mysql"
)

// staticly loaded data
var staticAnimals = []model.Animal{model.NewAnimal("Salsa", 4), model.NewAnimal("Wegorz", 2), model.NewAnimal("Krewetka", 5)}

var db *sql.DB

func GetAnimals(c *gin.Context) {

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
  fmt.Println("Connected!")

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

  c.IndentedJSON(http.StatusOK, animals)
}

func GetAnimalById(c *gin.Context) {
  id := c.Param("id")
  if id == "" {
    c.JSON(http.StatusBadRequest, gin.H{"errorMessage": "Missing ID path parameter"})
    return
  }

  for _, animal := range staticAnimals {
    if (animal.Id).String() == id {
      c.IndentedJSON(http.StatusOK, animal)
      return
    }
  }

  c.JSON(http.StatusNotFound, gin.H{"errorMessage": "Animal with given ID not found"})
}

func CreateAnimal(c *gin.Context) {

  var newAnimal model.Animal

  if err := c.BindJSON(&newAnimal); err != nil {
    return
  }

  staticAnimals = append(staticAnimals, newAnimal)
  c.IndentedJSON(http.StatusCreated, newAnimal)
}

func UpdateAnimal(c *gin.Context) {
  var newAnimal model.Animal

  fmt.Println(newAnimal)
  fmt.Println(c.BindJSON(&newAnimal))
  if err := c.BindJSON(&newAnimal); err != nil {
    fmt.Println("error ")
    c.JSON(http.StatusBadRequest, gin.H{"errorMessage": "Incorrect input data"})
    return
  }

  for i, animal := range staticAnimals {
    fmt.Println(animal)
    if animal.Id == newAnimal.Id {
      fmt.Println("weszlo")
      staticAnimals = append(staticAnimals[:i], staticAnimals[i+1:]...) // delete animal by slicing
      staticAnimals = append(staticAnimals, newAnimal) // add updated animal  TODO: investigate how to update it in one line

      c.IndentedJSON(http.StatusOK, newAnimal)
      return
    }
  }

  c.JSON(http.StatusNotFound, gin.H{"errorMessage": "Animal with given ID not found"})
}

func DeleteAnimalById(c *gin.Context) {
  id := c.Param("id")
  if id == "" {
     c.JSON(http.StatusBadRequest, gin.H{"errorMessage": "Missing ID in path"})
    return
  }

  for i, animal := range staticAnimals {
    if (animal.Id).String() == id {
      staticAnimals = append(staticAnimals[:i], staticAnimals[i+1:]...) // delete animal by slicing
      c.IndentedJSON(http.StatusOK, gin.H{"message": "Animal deleted"})
    }
  }
}
