// Code generated by MockGen. DO NOT EDIT.
// Source: animal-service.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	model "go-by-example/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAnimalService is a mock of AnimalService interface.
type MockAnimalService struct {
	ctrl     *gomock.Controller
	recorder *MockAnimalServiceMockRecorder
}

// MockAnimalServiceMockRecorder is the mock recorder for MockAnimalService.
type MockAnimalServiceMockRecorder struct {
	mock *MockAnimalService
}

// NewMockAnimalService creates a new mock instance.
func NewMockAnimalService(ctrl *gomock.Controller) *MockAnimalService {
	mock := &MockAnimalService{ctrl: ctrl}
	mock.recorder = &MockAnimalServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAnimalService) EXPECT() *MockAnimalServiceMockRecorder {
	return m.recorder
}

// AddAnimal mocks base method.
func (m *MockAnimalService) AddAnimal(newAnimal model.Animal) (model.Animal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAnimal", newAnimal)
	ret0, _ := ret[0].(model.Animal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAnimal indicates an expected call of AddAnimal.
func (mr *MockAnimalServiceMockRecorder) AddAnimal(newAnimal interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAnimal", reflect.TypeOf((*MockAnimalService)(nil).AddAnimal), newAnimal)
}

// DeleteAnimal mocks base method.
func (m *MockAnimalService) DeleteAnimal(id string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAnimal", id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAnimal indicates an expected call of DeleteAnimal.
func (mr *MockAnimalServiceMockRecorder) DeleteAnimal(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAnimal", reflect.TypeOf((*MockAnimalService)(nil).DeleteAnimal), id)
}

// GetAnimalById mocks base method.
func (m *MockAnimalService) GetAnimalById(id string) (model.Animal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAnimalById", id)
	ret0, _ := ret[0].(model.Animal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAnimalById indicates an expected call of GetAnimalById.
func (mr *MockAnimalServiceMockRecorder) GetAnimalById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAnimalById", reflect.TypeOf((*MockAnimalService)(nil).GetAnimalById), id)
}

// GetAnimals mocks base method.
func (m *MockAnimalService) GetAnimals() ([]model.Animal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAnimals")
	ret0, _ := ret[0].([]model.Animal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAnimals indicates an expected call of GetAnimals.
func (mr *MockAnimalServiceMockRecorder) GetAnimals() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAnimals", reflect.TypeOf((*MockAnimalService)(nil).GetAnimals))
}

// UpdateAnimal mocks base method.
func (m *MockAnimalService) UpdateAnimal(newAnimal model.Animal) (model.Animal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAnimal", newAnimal)
	ret0, _ := ret[0].(model.Animal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAnimal indicates an expected call of UpdateAnimal.
func (mr *MockAnimalServiceMockRecorder) UpdateAnimal(newAnimal interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAnimal", reflect.TypeOf((*MockAnimalService)(nil).UpdateAnimal), newAnimal)
}
