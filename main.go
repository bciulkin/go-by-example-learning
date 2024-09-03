package main

import (
  //"fmt"
  "net/http"
  "go-by-example/model"
  "encoding/json"
)

func main() {
  
  // static data
  animals := []model.Animal{model.NewAnimal("Salsa", 4), model.NewAnimal("Wegorz", 2), model.NewAnimal("Krewetka", 5)}

  http.HandleFunc("/animal", func (w http.ResponseWriter, r *http.Request) {
      
    if (r.Method == http.MethodPost) {
      decoder := json.NewDecoder(r.Body)
      var newAnimal model.Animal
      err := decoder.Decode(&newAnimal)
      if err != nil {
        panic(err)
      }

      animals = append(animals, newAnimal)
      w.WriteHeader(http.StatusCreated)
      json.NewEncoder(w).Encode(newAnimal)
    }

    if (r.Method == http.MethodPut) {
      decoder := json.NewDecoder(r.Body)
      var newAnimal model.Animal
      err := decoder.Decode(&newAnimal)
      if err != nil {
        panic(err)
      }

      for i, animal := range animals {
        if animal.Id == newAnimal.Id {
          animals = append(animals[:i], animals[i+1:]...) // delete animal by slicing
          animals = append(animals, newAnimal) // add updated animal  TODO: investigate how to update it in one line

          w.WriteHeader(http.StatusOK)
          json.NewEncoder(w).Encode(newAnimal)
          return
        }
      }
      http.Error(w, "Animal not found", http.StatusNotFound)
 
    }

    if (r.Method == http.MethodGet) {
      idStr := r.URL.Query().Get("id")
      if idStr == "" {
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(animals)
        return
      }

      for _, animal := range animals {
        if (animal.Id).String() == idStr {
          w.WriteHeader(http.StatusOK)
          json.NewEncoder(w).Encode(animal)
          return
        }
      }
      http.Error(w, "Animal not found", http.StatusNotFound)
    }

    if (r.Method == http.MethodDelete) {
      idStr := r.URL.Query().Get("id")
      if idStr == "" {
        http.Error(w, "missing id parameter", http.StatusBadRequest)
        return
      }

      for i, animal := range animals {
        if (animal.Id).String() == idStr {
          animals = append(animals[:i], animals[i+1:]...) // delete animal by slicing
          w.WriteHeader(http.StatusOK)
          json.NewEncoder(w).Encode(map[string]string{"message": "Animal deleted"})
          return
        }
      }
      http.Error(w, "Animal not found", http.StatusNotFound)
    }
  })

  http.ListenAndServe(":8080", nil)
}
