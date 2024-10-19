package main

import (
  "net/http"
  "log"
  //"time"
)

func main() {
  // define worker pool (5)
  
  
  p := NewPlayerParams()
  log.Println(p)
  http.PostForm("http://localhost:8080/player", p)
}

func worker(id int, jobs <-chan int, result chan<- int) {
  for j := range jobs {
    log.Println(j, id)
  //  p := NewPlayerParams()
    //resp, err := http.Post("http://localhost:8080/player", p)
  }
}
