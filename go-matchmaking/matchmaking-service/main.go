package main

import (
  "github.com/gin-gonic/gin"
 "net/http"
  "log"
)

func main() {
  router := gin.Default()
  router.GET("/pong", pong)

  log.Println("App is ready for ping-ponging.")
  router.Run("localhost:8080")
}

func pong(c *gin.Context) {

  c.IndentedJSON(http.StatusOK, "pong")
}
