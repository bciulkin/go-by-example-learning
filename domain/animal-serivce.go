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
