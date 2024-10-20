package main

import (
  "net/http"
  "fmt"
  "time"
)

func main() {
  // define worker pool (5)
  
  
  const numJobs = 10
  jobs := make(chan int, numJobs)

  for w := 1; w <= 4; w++ {
    go worker(w, jobs)
  }

  for j := 1; j <= numJobs; j++ {
    jobs <- j
  }
  close(jobs)
}

func worker(id int, jobs <-chan int) {
    for j := range jobs {
        time.Sleep(time.Second)
        http.PostForm("http://localhost:8080/player", NewPlayerParams())
        fmt.Println("Print, ", j)
    }
}
