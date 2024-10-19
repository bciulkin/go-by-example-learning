package main

import (
  "net/http"
  "log"
)

func main() {
  log.Println("Initialize ping-pong session.")
  resp, err := http.Post("http://localhost:8080/ping", "text/plain", "ping")
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()

}

