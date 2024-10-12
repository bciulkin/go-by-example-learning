package main

import (
  "log"
  "time"
  "sync"
)

func main() {
  log.Println("Tickers example:")
  
  ticker := time.NewTicker(500 * time.Millisecond)
  done := make(chan bool)

  go func() {
    for {
      select {
	case <-done:
	  return
	case <-ticker.C:
	  log.Println("Tick")
      }
    }
  }()

  time.Sleep(2100 * time.Millisecond)
  ticker.Stop()
  done <- true
  log.Println("Ticker example done.")

  // ************************************

	log.Println("Workers pool example:")
	const numJobs = 6
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 2; w++ {
		go worker(w, jobs, results)
	}

	for j :=1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a :=1; a <= numJobs; a++ {
		log.Println(<- results)
	}

	// ******************************
	// WaitGroup example
	var wg sync.WaitGroup
	for i := 1; i <=5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			workerI(i)
		}()
	}

	wg.Wait()
	log.Println("All done")
}

func workerI(id int) {
	log.Println("Worker ", id, "starting")
	time.Sleep(time.Second)
	log.Println("Worker ", id, "done")
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		log.Println("worker: ", id, "started job:", j)
		time.Sleep(time.Second)
		log.Println("worker: ", id, "started job:", j)
		results <- j*2
	}
}
