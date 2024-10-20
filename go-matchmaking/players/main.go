package main

import (
  "net/http"
  "fmt"
  "bytes"
  "io"
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
        url := "http://localhost:8080/player"

        var jsonStr = NewPlayerJsonString()
        //var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
        req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
        req.Header.Set("Content-Type", "application/json")
        
        client := &http.Client{}
        resp, err := client.Do(req)
        
        if err != nil {
          panic(err)
        }
        defer resp.Body.Close()

        fmt.Println("response Status:", resp.Status)
        fmt.Println("response Headers:", resp.Header)
        body, _ := io.ReadAll(resp.Body)
        fmt.Println("response Body:", string(body))

        results <- j * 2
    }
}
