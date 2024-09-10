package model

import (
  "fmt"
  "github.com/google/uuid"
)

type Animal struct {
  Id uuid.UUID  `json:"id"`
  Name string   `json:"name"`
  Age int       `json:"age"`
}

type Cat struct {
  Animal Animal
  CanMew bool
}

func NewAnimal(name string, age int) Animal {
  a := Animal{Id: uuid.New(), Name: name, Age: age}
  fmt.Println("Created animal with name: " +  a.Name)
  return a
}
