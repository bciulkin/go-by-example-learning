package main

import (
  "fmt"
  "net/http"
  "go-by-example/model"
  "encoding/json"
)

func main() {
  
  animal := model.NewAnimal("Salsa", 4)
  animalJson, _ := json.Marshal(animal)
  // cat := model.Cat{model.NewAnimal("Wegorz", 2), true}

  http.HandleFunc("/animal", func (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, string(animalJson))
  })

  http.ListenAndServe(":8080", nil)
}
