package domain

import (
  "testing"
  "go-by-example/model"
  "github.com/google/uuid"
  "github.com/stretchr/testify/assert"
  "github.com/DATA-DOG/go-sqlmock"
//  "fmt"
)

func TestAnimalRepository(t *testing.T) {
  db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
  if err != nil {
    t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
  }
  defer db.Close()

  repo := NewAnimalRepository(db)

  t.Run("getAllAnimals - successful case", func(t *testing.T) {
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
    rows := sqlmock.NewRows([]string{"name", "age", "id"}).
      AddRow("00000000-0000-0000-0000-000000000001","test", 1).
      AddRow("00000000-0000-0000-0000-000000000002","test2", 2)

    mock.ExpectQuery("SELECT * FROM animal").WillReturnRows(rows)

//    rs, err := db.Query("SELECT * FROM animal")
//    if err != nil {
//      fmt.Println("failed to match expected query")
//      return
//    }
//    defer rs.Close()

//    for rs.Next() {
//      var id uuid.UUID
//      var name string
//      var age int
//      rs.Scan(&name, &age, &id)
//      fmt.Println("scanned id:", id, "and name:", name)
//    }
//
//    if rs.Err() != nil {
//      fmt.Println("got rows error:", rs.Err())
//    }
    
    got, _ := repo.GetAllAnimals()

    assert.Equal(t, want, got)
  })
}
