package adapter

import (
  "testing"
  "go-by-example/model"
  "github.com/google/uuid"
  "github.com/golang/mock/gomock"
  mock "go-by-example/domain/mock"
//  "github.com/stretchr/testify/assert"
  "net/http/httptest"
  "net/http"
  "github.com/gin-gonic/gin"
//  "errors"
  "fmt"
)

func TestAnimalController(t *testing.T) {
  ctrl := gomock.NewController(t)

  service := mock.NewMockAnimalService(ctrl)
  controller := NewAnimalController(service)

  t.Run("getAnimalById - successful case", func(t *testing.T) {
    ctx := GetTestGinContext()

    want := model.Animal{
        Name: "test",
        Age: 1,
        Id: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
    }
    
    MockJsonGetAnimalById(ctx)
    service.EXPECT().GetAnimalById(gomock.Eq("00000000-0000-0000-0000-000000000001")).Return(want, nil).Times(1)

    controller.GetAnimalById(ctx)

    fmt.Println(ctx) // TODO: how to assert it
  })
}

func GetTestGinContext() *gin.Context {
    gin.SetMode(gin.TestMode)

    w := httptest.NewRecorder()
    ctx, _ := gin.CreateTestContext(w)
    ctx.Request = &http.Request{
      Header: make(http.Header),
    }

    return ctx
}

func MockJsonGetAnimalById(c *gin.Context) {
  c.Request.Method = "GET"
  c.Request.Header.Set("Content-Type", "application/json")
  c.Set("user_id", 1)

  // set path params
  c.Params = []gin.Param{
    {
      Key:   "id",
      Value: "00000000-0000-0000-0000-000000000001",
    },
  }
}
