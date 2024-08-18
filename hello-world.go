package main

import (
  "fmt"
  "net/http"
  "go-by-example/model"
)

func main() {
  
  fmt.Println(model.Animal{"Salsa", 4})
  fmt.Println(model.Cat{model.Animal{"Wegorz", 2}, true})

  http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
  })

  http.ListenAndServe(":8080", nil)
}
