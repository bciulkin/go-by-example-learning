package main

import (
  "github.com/gin-gonic/gin"
//  "net/http"
  "log"
)

func main() {
  router := gin.Default()
  router.POST("/ping", ping)

  log.Println("App is ready for ponging.")
  router.Run("localhost:8081")
}

func ping(c *gin.Context) {

}
