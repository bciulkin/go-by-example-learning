package main

import (
  "fmt"
  "encoding/json"
  "log"
  "os"
  "os/signal"

  "github.com/gorilla/websocket"
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
  	// Connect to the WebSocket server
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/player", nil)
	if err != nil {
		log.Fatal("Error connecting to WebSocket server:", err)
	}
	defer conn.Close()
        //var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
        
        // Send player info as JSON to the server
	playerData, _ := json.Marshal(NewPlayer())
	if err := conn.WriteMessage(websocket.TextMessage, playerData); err != nil {
		log.Println("Error sending player info:", err)
		return
	}
        
        // Channel to listen for interrupt signals (e.g., CTRL+C)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

        // Read messages from the server
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Error reading message:", err)
				return
			}

			// Check if it's a match or a waiting message
			var match Match
			if err := json.Unmarshal(message, &match); err == nil {
				fmt.Println("Match created! Here are the teams:")
				fmt.Println("Team 1:", match.Team1)
				fmt.Println("Team 2:", match.Team2)
			} else {
				// Print waiting message
				fmt.Println(string(message))
			}
		}
	}()

        // Keep the client running
        select {
	  case <-interrupt:
	    fmt.Println("Client interrupted, closing connection.")
	}

        results <- j * 2
    }
}


