package domain

import (
  "go-by-example/model"
)

//go:generate mockgen -source=animal-service.go -destination=mock/animal-service.go

type AnimalService interface {
  GetAnimals() ([]model.Animal, error)
  GetAnimalById(id string) (model.Animal, error)
  AddAnimal(newAnimal model.Animal) (model.Animal, error)
  UpdateAnimal(newAnimal model.Animal) (model.Animal, error)
  DeleteAnimal(id string) (string, error)
}

type animalService struct {
  AnimalRepository AnimalRepository
}

func NewAnimalService(animalRepository AnimalRepository) AnimalService {
  return &animalService{
    AnimalRepository: animalRepository,
  }
}

func (service *animalService) GetAnimals() ([]model.Animal, error) {
  animals, err := service.AnimalRepository.GetAllAnimals()
  if err != nil {
    return animals, err
  }

  return animals, nil
}

func (service *animalService) GetAnimalById(id string) (model.Animal, error) {
  animal, err := service.AnimalRepository.GetAnimalById(id)
  if err != nil {
    return animal, err
  }
  return animal, nil
}

func (service *animalService) AddAnimal(newAnimal model.Animal) (model.Animal, error) {

  createdAnimal, err := service.AnimalRepository.AddAnimal(newAnimal)
  if err != nil {
    return createdAnimal, err
  }
  return createdAnimal, nil
}

func (service *animalService) UpdateAnimal(newAnimal model.Animal) (model.Animal, error) {

  updatedAnimal, err := service.AnimalRepository.UpdateAnimal(newAnimal)
  if err != nil {
    return updatedAnimal, err
  }
  return updatedAnimal, nil
}

func (service *animalService) DeleteAnimal(id string) (string, error) {
  _, err := service.AnimalRepository.DeleteAnimal(id)
  if err != nil {
    return id, err
  }

  return id, nil
}
