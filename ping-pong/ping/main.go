package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "log"
  "time"
)

func main() {
  router := gin.Default()
  router.POST("/ping", ping)

  log.Println("App is ready for ping-pong-ing.")
  router.Run("localhost:8081")
}

func ping(c *gin.Context) {
  resp, err := http.Get("http://localhost:8080/pong")
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()

  log.Println(resp.Status)
  if resp.Status == "200 OK" {
    time.Sleep(10 * time.Second)
    ping(c)
    
    c.IndentedJSON(http.StatusOK, "pong")
  }
}
