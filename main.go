package main

import (
  "github.com/gin-gonic/gin"
  "go-by-example/adapter"
  "go-by-example/domain"
  "log"
)

func main() {
  domain.ConnectToDb()

  router := gin.Default()
  router.GET("/animal", controller.GetAnimals)
  router.GET("/animal/:id", controller.GetAnimalById)

  router.POST("/animal", controller.CreateAnimal)
  router.DELETE("/animal/:id", controller.DeleteAnimalById)
  router.PUT("/animal", controller.UpdateAnimal)

  log.Println("App has started. About to start serving REST API.")
  router.Run("localhost:8080")
}
