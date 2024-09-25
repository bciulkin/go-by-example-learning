package domain

import (
  "testing"
  "go-by-example/model"
  "github.com/google/uuid"
  "github.com/golang/mock/gomock"
  mock "go-by-example/domain/mock"
  "github.com/stretchr/testify/assert"
)
// TODO: split to t.Run cases
// TODO: add more cases

func TestGetAnimals(t *testing.T) {
  ctrl := gomock.NewController(t)

  repo := mock.NewMockAnimalRepository(ctrl)
  service := NewAnimalService(repo)

  // mock return from repository
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
  repo.EXPECT().GetAllAnimals().Return(want, nil).Times(1)

  got, err := service.GetAnimals()
  if err != nil {
    t.Errorf("service.GetAnimals error")
  }

  assert.Equal(t, got, want)
}
