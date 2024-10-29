package main

import (
  "fmt"
  "encoding/json"
  "log"
  "os"
  "os/signal"
  "github.com/gorilla/websocket"
  "sync"
  "time"
  "math/rand"
)

func main() {
	fmt.Println("*** Populate players ***")
	rand.Seed(time.Now().UnixNano())

	// Set up the WorkerPool with the desired number of concurrent workers
	pool := NewWorkerPool(5)
	pool.Start()

	// Signal channel to handle CTRL+C for graceful shutdown
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Generate random player jobs and add them to the pool
	go func() {
		for i := 1; i <= 20; i++ {
			player := NewPlayer()
			pool.AddJob(player)
			time.Sleep(200 * time.Millisecond) // Adjust to control the rate of new jobs
		}
	}()

	// Wait for an interrupt signal to stop the pool
	<-interrupt
	fmt.Println("\nStopping worker pool...")
	pool.Stop()
	fmt.Println("Worker pool stopped.")
}

type WorkerPool struct {
  numWorkers int
  jobs       chan Player
  wg         sync.WaitGroup
}

func NewWorkerPool(numWorkers int) *WorkerPool {
	return &WorkerPool{
		numWorkers: numWorkers,
		jobs:       make(chan Player, numWorkers),
	}
}

// Start initializes workers and assigns them jobs
func (wp *WorkerPool) Start() {
	for i := 0; i < wp.numWorkers; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

// Stop waits for all workers to complete
func (wp *WorkerPool) Stop() {
	close(wp.jobs)
	wp.wg.Wait()
}

func (wp *WorkerPool) AddJob(player Player) {
	wp.jobs <- player
}

func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()

	for player := range wp.jobs {
		fmt.Printf("worker %d connecting to server for player: %s\n", id, player.Id)

		conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/player", nil)
		if err != nil {
			log.Fatal("Error connecting to WebSocket server:", err)
		}
		defer conn.Close()

		// Send player info as JSON to the server
		playerData, _ := json.Marshal(player)
		if err := conn.WriteMessage(websocket.TextMessage, playerData); err != nil {
			log.Println("Error sending player info:", err)
			continue
		}
  
		// Listen for messages from the server
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Worker %d: Error reading message: %v", id, err)
			continue
		}

		// Display the server response
		fmt.Printf("Worker %d received: %s\n", id, message)
	}
}


