package service

import (
	"errors"
	"net/http/httptest"
	"testing"

	entity "CrudPlatform/internal/core/domain/repository"
	model "CrudPlatform/internal/core/domain/repository/model/users"
	schema "CrudPlatform/internal/core/domain/repository/schema/users"
	mockRepository "CrudPlatform/internal/core/ports/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewService(t *testing.T) {
	mockRepo := mockRepository.NewDBRepositoryUsers(t)
	service := NewService(mockRepo)
	assert.NotNil(t, service, "El servicio no debe ser nil")
}

func TestCreateUser(t *testing.T) {
	mockRepo := mockRepository.NewDBRepositoryUsers(t)
	svc := &Repository{repo: mockRepo}

	mockRepo.On("CreateUser", mock.Anything, mock.Anything).Return("123", nil)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := &model.User{
		Name:      "John Doe",
		Email:     "john@example.com",
		ImagePath: "/path/to/image.jpg",
	}

	expectedResp := &entity.Response{
		Data: "123",
		Result: entity.Result{
			Details: []entity.Detail{
				{InternalCode: "200", Message: "OK", Detail: "Registro Creado"},
			},
			Source: "Create User",
		},
	}

	response, err := svc.CreateUser(c, req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp, response)
}

func TestCreateUser_ErrorCase(t *testing.T) {
	mockRepo := mockRepository.NewDBRepositoryUsers(t)
	svc := &Repository{repo: mockRepo}

	mockRepo.On("CreateUser", mock.Anything, mock.Anything).Return("", errors.New("error simulado"))

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := &model.User{
		Name:      "John Doe",
		Email:     "john@example.com",
		ImagePath: "/path/to/image.jpg",
	}

	response, err := svc.CreateUser(c, req)
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestSelectUser(t *testing.T) {
	mockRepo := mockRepository.NewDBRepositoryUsers(t)
	svc := &Repository{repo: mockRepo}

	mockResponse := &schema.UsersGetResponse{
		Name:      "John Doe",
		Email:     "john@example.com",
		ImagePath: "/path/to/image.jpg",
	}
	mockRepo.On("SelectUser", mock.Anything, mock.Anything).Return(mockResponse, nil)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := &model.GetUser{Id: "123"}

	expectedResp := &entity.Response{
		Data: mockResponse,
		Result: entity.Result{
			Details: []entity.Detail{
				{InternalCode: "200", Message: "OK", Detail: "Registro Seleccionado"},
			},
			Source: "Select User",
		},
	}

	response, err := svc.SelectUser(c, req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp, response)
}

func TestSelectUser_ErrorCase(t *testing.T) {
	mockRepo := mockRepository.NewDBRepositoryUsers(t)
	svc := &Repository{repo: mockRepo}

	mockRepo.On("SelectUser", mock.Anything, mock.Anything).Return(nil, errors.New("error simulado"))

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := &model.GetUser{Id: "123"}

	response, err := svc.SelectUser(c, req)
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestUpdateUser(t *testing.T) {
	mockRepo := mockRepository.NewDBRepositoryUsers(t)
	svc := &Repository{repo: mockRepo}

	mockResponse := &schema.UsersUpdateResponse{
		Name:      "John Doe Updated",
		Email:     "john_updated@example.com",
		ImagePath: "/path/to/new_image.jpg",
	}
	mockRepo.On("UpdateUser", mock.Anything, mock.Anything).Return(mockResponse, nil)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := &model.UpdateUser{
		Id:        "123",
		Name:      "John Doe Updated",
		Email:     "john_updated@example.com",
		ImagePath: "/path/to/new_image.jpg",
	}

	expectedResp := &entity.Response{
		Data: mockResponse,
		Result: entity.Result{
			Details: []entity.Detail{
				{InternalCode: "200", Message: "OK", Detail: "Registro Actualizado"},
			},
			Source: "Update User",
		},
	}

	response, err := svc.UpdateUser(c, req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp, response)
}

func TestUpdateUser_ErrorCase(t *testing.T) {
	mockRepo := mockRepository.NewDBRepositoryUsers(t)
	svc := &Repository{repo: mockRepo}

	mockRepo.On("UpdateUser", mock.Anything, mock.Anything).Return(nil, errors.New("error simulado"))

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := &model.UpdateUser{
		Id:        "123",
		Name:      "John Doe Updated",
		Email:     "john_updated@example.com",
		ImagePath: "/path/to/new_image.jpg",
	}

	response, err := svc.UpdateUser(c, req)
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestDeleteUser(t *testing.T) {
	mockRepo := mockRepository.NewDBRepositoryUsers(t)
	svc := &Repository{repo: mockRepo}

	mockRepo.On("DeleteUser", mock.Anything, mock.Anything).Return(nil)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := &model.DeleteUser{Id: "123"}

	expectedResp := &entity.Response{
		Result: entity.Result{
			Details: []entity.Detail{
				{InternalCode: "200", Message: "OK", Detail: "Registro Eliminado"},
			},
			Source: "Delete User",
		},
	}

	response, err := svc.DeleteUser(c, req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp, response)
}

func TestDeleteUser_ErrorCase(t *testing.T) {
	mockRepo := mockRepository.NewDBRepositoryUsers(t)
	svc := &Repository{repo: mockRepo}

	mockRepo.On("DeleteUser", mock.Anything, mock.Anything).Return(errors.New("error simulado"))

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := &model.DeleteUser{Id: "123"}

	response, err := svc.DeleteUser(c, req)
	assert.Error(t, err)
	assert.Nil(t, response)
}
