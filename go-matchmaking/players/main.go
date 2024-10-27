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
  
  for w := 1; w <= 10; w++ {
    go worker(w, NewPlayer())
  }

}

func worker(id int, player Player) {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/player", nil)
	if err != nil {
		log.Fatal("Error connecting to WebSocket server:", err)
	}
	defer conn.Close()
        
        // Send player info as JSON to the server
	playerData, _ := json.Marshal(player)
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
}


