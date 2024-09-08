package domain

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "net/http"
  "go-by-example/model"
)

// staticly loaded data
var animals = []model.Animal{model.NewAnimal("Salsa", 4), model.NewAnimal("Wegorz", 2), model.NewAnimal("Krewetka", 5)}

func GetAnimals(c *gin.Context) {
  c.IndentedJSON(http.StatusOK, animals)
}

func GetAnimalById(c *gin.Context) {
  id := c.Param("id")
  if id == "" {
    c.JSON(http.StatusBadRequest, gin.H{"errorMessage": "Missing ID path parameter"})
    return
  }

  for _, animal := range animals {
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

  animals = append(animals, newAnimal)
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

  for i, animal := range animals {
    fmt.Println(animal)
    if animal.Id == newAnimal.Id {
      fmt.Println("weszlo")
      animals = append(animals[:i], animals[i+1:]...) // delete animal by slicing
      animals = append(animals, newAnimal) // add updated animal  TODO: investigate how to update it in one line

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

  for i, animal := range animals {
    if (animal.Id).String() == id {
      animals = append(animals[:i], animals[i+1:]...) // delete animal by slicing
      c.IndentedJSON(http.StatusOK, gin.H{"message": "Animal deleted"})
    }
  }
}
