package domain

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "net/http"
  "go-by-example/model"
)

// staticly loaded data
var staticAnimals = []model.Animal{model.NewAnimal("Salsa", 4), model.NewAnimal("Wegorz", 2), model.NewAnimal("Krewetka", 5)}

func GetAnimals(c *gin.Context) {

  animals := getAllAnimals()
  c.IndentedJSON(http.StatusOK, animals)
}

func GetAnimalById(c *gin.Context) {
  id := c.Param("id")
  if id == "" {
    c.JSON(http.StatusBadRequest, gin.H{"errorMessage": "Missing ID path parameter"})
    return
  }

  animal, err := getAnimalById(id)
  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{"errorMessage": "Animal with given ID not found"})
  }

  c.IndentedJSON(http.StatusOK, animal)
}

func CreateAnimal(c *gin.Context) {

  var newAnimal model.Animal

  if err := c.BindJSON(&newAnimal); err != nil {
    return
  }

  createdAnimal, err := addAnimal(newAnimal)
  if err != nil {
    c.IndentedJSON(http.StatusBadRequest, gin.H{"errorMessage": err})
  }

  c.IndentedJSON(http.StatusCreated, createdAnimal)
}

func UpdateAnimal(c *gin.Context) {
  var newAnimal model.Animal

  if jsonErr := c.BindJSON(&newAnimal); jsonErr != nil {
    fmt.Println("error ")
    c.JSON(http.StatusBadRequest, gin.H{"errorMessage": "Incorrect input data"})
    return
  }
  
  updatedAnimal, err := updateAnimal(newAnimal)
  if err != nil {
    c.IndentedJSON(http.StatusBadRequest, gin.H{"errorMessage": err})
  }

  c.IndentedJSON(http.StatusOK, updatedAnimal)
}

func DeleteAnimalById(c *gin.Context) {
  id := c.Param("id")
  if id == "" {
     c.JSON(http.StatusBadRequest, gin.H{"errorMessage": "Missing ID in path"})
    return
  }
  _, err := deleteAnimal(id)
  if err != nil {
    c.IndentedJSON(http.StatusBadRequest, gin.H{"errorMessage": err})
  }

  c.IndentedJSON(http.StatusOK, gin.H{"message": "Animal deleted"})
}
