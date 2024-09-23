package domain

import (
  "log"
  "database/sql"
  "go-by-example/model"
  "fmt"
)

type AnimalRepository interface {
  GetAllAnimals() ([]model.Animal, error)
  GetAnimalById(id string) (model.Animal, error)
  AddAnimal(anml model.Animal) (model.Animal, error)
  UpdateAnimal(anml model.Animal) (model.Animal, error)
  DeleteAnimal(id string) (string, error)
}

type animalRepository struct {
}

func NewAnimalRepository() AnimalRepository {
  return &animalRepository{}
}

func (repository *animalRepository) GetAllAnimals() ([]model.Animal, error) {
  var animals []model.Animal
  rows, err := Db.Query("SELECT * FROM animal")
  if err != nil {
    return animals, fmt.Errorf("getAllAnimals: %v", err)
  }
  defer rows.Close()
  for rows.Next() {
    var animal model.Animal
    if err := rows.Scan(&animal.Id, &animal.Name, &animal.Age); err != nil {
      return animals, fmt.Errorf("getAllAnimals: %v", err)
    }
    animals = append(animals, animal)
  }
  return animals, nil
}

func (repository *animalRepository) GetAnimalById(id string) (model.Animal, error) {
  var animal model.Animal

  row := Db.QueryRow("SELECT * FROM animal WHERE id = ?", id)
  if err := row.Scan(&animal.Id, &animal.Name, &animal.Age); err != nil {
    if err == sql.ErrNoRows {
      return animal, fmt.Errorf("Animal with ID not found: %s", id)
    } else {
      log.Println(err)
    }
  }
  return animal, nil
}

func (repository *animalRepository) AddAnimal(anml model.Animal) (model.Animal, error) {
  _, err := Db.Exec("INSERT INTO animal (id, name, age) VALUES (?, ?, ?)", anml.Id, anml.Name, anml.Age)
  if err != nil {
    return anml, fmt.Errorf("addAnimal: %v", err)
  }
  return repository.GetAnimalById(anml.Id.String())
  
}

func (repository *animalRepository) UpdateAnimal(anml model.Animal) (model.Animal, error) {
  _, err := Db.Exec("UPDATE animal SET name = ?, age = ? WHERE id = ?", anml.Name, anml.Age, anml.Id)
  if err != nil {
    return anml, fmt.Errorf("updateAnimal: %v", err)
  }
  return repository.GetAnimalById(anml.Id.String())
}

func (repository *animalRepository) DeleteAnimal(id string) (string, error) {
  _, err := Db.Exec("DELETE FROM animal WHERE id = ?", id)
  if err != nil {
    return id, fmt.Errorf("deleteAnimal: %v", err)
  }
  return id, nil
}
