package main

import (
  "fmt"
  "net/http"
  "go-by-example/model"
)

func main() {
  
  animal := model.NewAnimal("Salsa", 4)
  // cat := model.Cat{model.NewAnimal("Wegorz", 2), true}

  http.HandleFunc("/animal", func (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Name: %v, age: %v", animal.Name, animal.Age)
  })

  http.ListenAndServe(":8080", nil)
}
