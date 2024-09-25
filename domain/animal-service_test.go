package domain

import (
  "testing"
  "go-by-example/model"
  "github.com/google/uuid"
  "github.com/golang/mock/gomock"
  mock "go-by-example/domain/mock"
  "github.com/stretchr/testify/assert"
  "errors"
)
// TODO: split to t.Run cases
// TODO: add more cases

func TestGetAnimals(t *testing.T) {
  ctrl := gomock.NewController(t)

  repo := mock.NewMockAnimalRepository(ctrl)
  service := NewAnimalService(repo)

  t.Run("successful case", func(t *testing.T) {
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

    got, _ := service.GetAnimals()

    assert.Equal(t, got, want)
  })

  t.Run("error case", func(t *testing.T) {
    want := errors.New("test error")
    repo.EXPECT().GetAllAnimals().Return(nil, want).Times(1)

    _, err := service.GetAnimals()

    assert.ErrorIs(t, err, want)
  })
}
