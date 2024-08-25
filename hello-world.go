package main

import (
  "fmt"
  "net/http"
  "go-by-example/model"
  "encoding/json"
)

func main() {
  
  // cat := model.Cat{model.NewAnimal("Wegorz", 2), true}

  // GET /animal
  http.HandleFunc("/animal", func (w http.ResponseWriter, r *http.Request) {
    var animal model.Animal = model.NewAnimal("Salsa", 4)
    var animalJson, _ = json.Marshal(animal)
    fmt.Fprintf(w, string(animalJson))
  })

  // POST /test
  http.HandleFunc("/test", func (rw http.ResponseWriter, req *http.Request) {
    decoder := json.NewDecoder(req.Body)
    var a model.Animal
    err := decoder.Decode(&a)
    if err != nil {
        panic(err)
    }

    aJosn, _ := json.Marshal(a)
    fmt.Fprintf(rw, string(aJosn))
  })

  http.ListenAndServe(":8080", nil)
}
