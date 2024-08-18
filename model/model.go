package model

type Animal struct {
  Name string
  Age int
}

type Cat struct {
  Animal Animal
  CanMew bool
}

func newAnimal(name string) *Animal {
  a := Animal{Name: name}
  a.Age = 3
  return &a
}
