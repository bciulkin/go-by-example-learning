package domain

import (
  "testing"
  "go-by-example/model"
  "github.com/google/uuid"
  "github.com/google/go-cmp/cmp"
)

func TestGetAnimals(t *testing.T) {
  repo := NewAnimalRepository()
  service := NewAnimalService(repo)

  got, err := service.GetAnimals()
  if err != nil {
    t.Errorf("service.GetAnimals error")
  }
  want := []model.Animal{
    {
      Name: "test",
      Age: 1,
      Id: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
    },
    {
      Name: "test2",
      Age: 2,
      Id: uuid.MustParse("00000000-0000-0000-0000-000000000002"),
    },
  }

  if cmp.Equal(got, want) {
    t.Errorf("got %q wanted %q", got, want)
  }
}
