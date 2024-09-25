package main

import (
  "github.com/gin-gonic/gin"
  "go-by-example/adapter"
  "go-by-example/domain"
  "log"
  "os"
)

func main() {
  args := os.Args[1:]
  dbUser := args[0]
  dbPass := args[1]

  db := domain.ConnectToDb(dbUser, dbPass)

  repository := domain.NewAnimalRepository(db)
  service := domain.NewAnimalService(repository)
  controller := adapter.NewAnimalController(service)

  router := gin.Default()
  router.GET("/animal", controller.GetAnimals)
  router.GET("/animal/:id", controller.GetAnimalById)

  router.POST("/animal", controller.CreateAnimal)
  router.DELETE("/animal/:id", controller.DeleteAnimalById)
  router.PUT("/animal", controller.UpdateAnimal)

  log.Println("App has started. About to start serving REST API.")
  router.Run("localhost:8080")
}
