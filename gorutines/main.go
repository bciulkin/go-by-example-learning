package main

import (
  "log"
  "time"
)

func main() {
  // one goroutine
  message := make(chan string)

  go func() {message <- "ping"}()

  msg := <-message

  log.Println(msg)

  // **********************************

  // channeled two goroutines
  twoMessages := make(chan string, 2)

  twoMessages <- "one"
  twoMessages <- "two"

  log.Println(<-twoMessages)
  log.Println(<-twoMessages)

  // **********************************

  // syncing + method definition with channel (worker)
  done := make(chan bool, 2)
  go worker(done)
  go worker(done)

  log.Println("first, ", <-done)
  log.Println("second, ", <-done)

  // **********************************
  // channel direction ping pong

  pings := make(chan string, 1)
  pongs := make(chan string, 1)
  ping(pings, "one way ticket")
  pong(pings, pongs)
  log.Println(<-pongs)


  // **********************************
  // select feature

  c1 := make(chan string)
  c2 := make(chan string)

  go func() {
    time.Sleep(1 * time.Second)
    c1 <- "one"
  }()
  go func() {
    time.Sleep(2 * time.Second)
    c2 <- "two"
  }()

  for i := 0; i < 2; i++ {
    select {
      case msg1 := <-c1:
        log.Println("received", msg1)
      case msg2 := <-c2:
        log.Println("received", msg2)
    }
  }
}

func worker(done chan bool) {

  log.Println("working..")
  time.Sleep(time.Second)
  log.Println("done")
  done <-true
}

func ping(pings chan <- string, message string) {
  pings <- message
}

func pong(pings <-chan string, pongs chan<- string) {
  msg := <-pings
  log.Println("on the way: ", msg)
  pongs <- msg
}
