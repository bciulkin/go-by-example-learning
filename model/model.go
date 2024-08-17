package model

type Animal struct {
  Name string
  Age int
}

func newAnimal(name string) *Animal {
  a := Animal{Name: name}
  a.Age = 3
  return &a
}
