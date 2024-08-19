package model

import (
  "fmt"
  "github.com/google/uuid"
)

type Animal struct {
  id uuid.UUID
  Name string
  Age int
}

type Cat struct {
  Animal Animal
  CanMew bool
}

func NewAnimal(name string, age int) Animal {
  a := Animal{id: uuid.New(), Name: name, Age: age}
  fmt.Println(a.id)
  return a
}
