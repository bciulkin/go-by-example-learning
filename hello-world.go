package main

import (
  "fmt"
  "net/http"
  "go-by-example/model"
  "encoding/json"
)

func main() {
  
  // cat := model.Cat{model.NewAnimal("Wegorz", 2), true}
  
  // static data
  animals := []model.Animal{model.NewAnimal("Salsa", 4), model.NewAnimal("Wegorz", 2)}

  // GET /animal
  http.HandleFunc("/animal", func (w http.ResponseWriter, r *http.Request) {
    if (r.Method == http.MethodGet) {
      var animalJson, _ = json.Marshal(animals[0])
      fmt.Fprintf(w, string(animalJson))
    }

    if (r.Method == http.MethodDelete) {
      idStr := r.URL.Query().Get("id")
      if idStr == "" {
        http.Error(w, "missing id parameter", http.StatusBadRequest)
        return
      }

      // Respond with success
      for i, animal := range animals {
        if (animal.Id).String() == idStr {
          animals = append(animals[:i], animals[i+1:]...) // delete animal by slicing
          w.WriteHeader(http.StatusOK)
          json.NewEncoder(w).Encode(map[string]string{"message": "Animal deleted"})
          return
        }
      }

      // TODO: delete handling after introducing SQLite

      http.Error(w, "Animal not found", http.StatusNotFound)
    }
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
