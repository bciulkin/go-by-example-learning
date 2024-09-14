package domain

import (
  "go-by-example/model"
)

func GetAnimals() ([]model.Animal, error) {
  animals, err := getAllAnimals()
  if err != nil {
    return animals, err
  }

  return animals, nil
}

func GetAnimalById(id string) (model.Animal, error) {
  animal, err := getAnimalById(id)
  if err != nil {
    return animal, err
  }
  return animal, nil
}

func AddAnimal(newAnimal model.Animal) (model.Animal, error) {

  createdAnimal, err := addAnimal(newAnimal)
  if err != nil {
    return createdAnimal, err
  }
  return createdAnimal, nil
}

func UpdateAnimal(newAnimal model.Animal) (model.Animal, error) {

  updatedAnimal, err := updateAnimal(newAnimal)
  if err != nil {
    return updatedAnimal, err
  }
  return updatedAnimal, nil
}

func DeleteAnimal(id string) (string, error) {
  _, err := deleteAnimal(id)
  if err != nil {
    return id, err
  }

  return id, nil
}
