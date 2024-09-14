package controller

import (
  "log"
  "github.com/gin-gonic/gin"
  "net/http"
  "go-by-example/model"
  "go-by-example/domain"
)

func GetAnimals(c *gin.Context) {
  animals, err := domain.GetAnimals()
  if err != nil {
    c.IndentedJSON(http.StatusBadRequest, gin.H{"errorMessage": err})
    return
  }

  c.IndentedJSON(http.StatusOK, animals)
}

func GetAnimalById(c *gin.Context) {
  id := c.Param("id")
  if id == "" {
    c.IndentedJSON(http.StatusBadRequest, gin.H{"errorMessage": "Missing ID path parameter"})
    return
  }

  animal, err := domain.GetAnimalById(id)
  if err != nil {
    c.IndentedJSON(http.StatusNotFound, gin.H{"errorMessage": "Animal with given ID not found"})
    return
  }

  c.IndentedJSON(http.StatusOK, animal)
}

func CreateAnimal(c *gin.Context) {

  var newAnimal model.Animal

  if err := c.BindJSON(&newAnimal); err != nil {
    return
  }

  createdAnimal, err := domain.AddAnimal(newAnimal)
  if err != nil {
    c.IndentedJSON(http.StatusBadRequest, gin.H{"errorMessage": err})
    return
  }

  c.IndentedJSON(http.StatusCreated, createdAnimal)
}

func UpdateAnimal(c *gin.Context) {
  var newAnimal model.Animal

  if jsonErr := c.BindJSON(&newAnimal); jsonErr != nil {
    log.Println("error")
    c.IndentedJSON(http.StatusBadRequest, gin.H{"errorMessage": "Incorrect input data"})
    return
  }
  
  updatedAnimal, err := domain.UpdateAnimal(newAnimal)
  if err != nil {
    c.IndentedJSON(http.StatusBadRequest, gin.H{"errorMessage": err})
    return
  }

  c.IndentedJSON(http.StatusOK, updatedAnimal)
}

func DeleteAnimalById(c *gin.Context) {
  id := c.Param("id")
  if id == "" {
     c.IndentedJSON(http.StatusBadRequest, gin.H{"errorMessage": "Missing ID in path"})
     return
  }
  _, err := domain.DeleteAnimal(id)
  if err != nil {
    c.IndentedJSON(http.StatusBadRequest, gin.H{"errorMessage": err})
    return
  }

  c.IndentedJSON(http.StatusOK, gin.H{"message": "Animal deleted"})
}
