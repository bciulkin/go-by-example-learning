package main

import (
  "net/http"
  "fmt"
)

func main() {
  // define worker pool (5)
  
  
  const numJobs = 10
  jobs := make(chan int, numJobs)
  results := make(chan int, numJobs)

  for w := 1; w <= 4; w++ {
    go worker(w, jobs, results)
  }

  for j := 1; j <= numJobs; j++ {
    jobs <- j
  }
  close(jobs)

  for a := 1; a <= numJobs; a++ {
    <-results
  }
}

func worker(id int, jobs <-chan int, results chan<- int) {

    fmt.Println("Print, worker")
    for j := range jobs {
        fmt.Println("Print, ", j, id)
        http.PostForm("http://localhost:8080/player", NewPlayerParams())
        results <- j * 2
    }
}
