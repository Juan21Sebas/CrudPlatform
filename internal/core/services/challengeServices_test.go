package service

import (
	"errors"
	"net/http/httptest"
	"testing"

	entity "CrudPlatform/internal/core/domain/repository"
	model "CrudPlatform/internal/core/domain/repository/model/challenges"
	schema "CrudPlatform/internal/core/domain/repository/schema/challenges"
	mockRepository "CrudPlatform/internal/core/ports/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewServiceChallenge(t *testing.T) {
	mockRepo := mockRepository.NewDBRepositoryChallenge(t)
	service := NewServiceChallenge(mockRepo)
	assert.NotNil(t, service, "El servicio no debe ser nil")
}

func TestCreateChallenge(t *testing.T) {
	mockRepo := mockRepository.NewDBRepositoryChallenge(t)
	svc := &RepositoryChallenge{repo: mockRepo}

	mockRepo.On("CreateChallenge", mock.Anything, mock.Anything).Return("123", nil)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := &model.Challenge{
		Title:       "Test Challenge",
		Description: "Test Description",
		Difficulty:  3,
	}

	expectedResp := &entity.Response{
		Data: "123",
		Result: entity.Result{
			Details: []entity.Detail{
				{InternalCode: "200", Message: "OK", Detail: "Registro Creado"},
			},
			Source: "Create Challenge",
		},
	}

	response, err := svc.CreateChallenge(c, req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp, response)
}

func TestCreateChallenge_ErrorCase(t *testing.T) {
	mockRepo := mockRepository.NewDBRepositoryChallenge(t)
	svc := &RepositoryChallenge{repo: mockRepo}

	mockRepo.On("CreateChallenge", mock.Anything, mock.Anything).Return("", errors.New("error simulado"))

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := &model.Challenge{
		Title:       "Test Challenge",
		Description: "Test Description",
		Difficulty:  3,
	}

	response, err := svc.CreateChallenge(c, req)
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestSelectChallenge(t *testing.T) {
	mockRepo := mockRepository.NewDBRepositoryChallenge(t)
	svc := &RepositoryChallenge{repo: mockRepo}

	mockResponse := &schema.ChallengeGetResponse{
		Title:       "Test Challenge",
		Description: "Test Description",
		Difficulty:  3,
	}
	mockRepo.On("SelectChallenge", mock.Anything, mock.Anything).Return(mockResponse, nil)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := &model.GetChallenge{ID: "123"}

	expectedResp := &entity.Response{
		Data: mockResponse,
		Result: entity.Result{
			Details: []entity.Detail{
				{InternalCode: "200", Message: "OK", Detail: "Registro Seleccionado"},
			},
			Source: "Select Challenge",
		},
	}

	response, err := svc.SelectChallenge(c, req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp, response)
}

func TestSelectChallenge_ErrorCase(t *testing.T) {
	mockRepo := mockRepository.NewDBRepositoryChallenge(t)
	svc := &RepositoryChallenge{repo: mockRepo}

	mockRepo.On("SelectChallenge", mock.Anything, mock.Anything).Return(nil, errors.New("error simulado"))

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := &model.GetChallenge{ID: "123"}

	response, err := svc.SelectChallenge(c, req)
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestUpdateChallenge(t *testing.T) {
	mockRepo := mockRepository.NewDBRepositoryChallenge(t)
	svc := &RepositoryChallenge{repo: mockRepo}

	mockResponse := &schema.ChallengeUpdateResponse{
		Title:       "Updated Challenge",
		Description: "Updated Description",
		Difficulty:  4,
	}
	mockRepo.On("UpdateChallenge", mock.Anything, mock.Anything).Return(mockResponse, nil)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := &model.UpdateChallenge{
		ID:          "123",
		Title:       "Updated Challenge",
		Description: "Updated Description",
		Difficulty:  4,
	}

	expectedResp := &entity.Response{
		Data: mockResponse,
		Result: entity.Result{
			Details: []entity.Detail{
				{InternalCode: "200", Message: "OK", Detail: "Registro Actualizado"},
			},
			Source: "Update Challenge",
		},
	}

	response, err := svc.UpdateChallenge(c, req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp, response)
}

func TestUpdateChallenge_ErrorCase(t *testing.T) {
	mockRepo := mockRepository.NewDBRepositoryChallenge(t)
	svc := &RepositoryChallenge{repo: mockRepo}

	mockRepo.On("UpdateChallenge", mock.Anything, mock.Anything).Return(nil, errors.New("error simulado"))

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := &model.UpdateChallenge{
		ID:          "123",
		Title:       "Updated Challenge",
		Description: "Updated Description",
		Difficulty:  4,
	}

	response, err := svc.UpdateChallenge(c, req)
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestDeleteChallenge(t *testing.T) {
	mockRepo := mockRepository.NewDBRepositoryChallenge(t)
	svc := &RepositoryChallenge{repo: mockRepo}

	mockRepo.On("DeleteChallenge", mock.Anything, mock.Anything).Return(nil)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := &model.DeleteChallenge{ID: "123"}

	expectedResp := &entity.Response{
		Result: entity.Result{
			Details: []entity.Detail{
				{InternalCode: "200", Message: "OK", Detail: "Registro Eliminado"},
			},
			Source: "Delete Challenge",
		},
	}

	response, err := svc.DeleteChallenge(c, req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp, response)
}

func TestDeleteChallenge_ErrorCase(t *testing.T) {
	mockRepo := mockRepository.NewDBRepositoryChallenge(t)
	svc := &RepositoryChallenge{repo: mockRepo}

	mockRepo.On("DeleteChallenge", mock.Anything, mock.Anything).Return(errors.New("error simulado"))

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := &model.DeleteChallenge{ID: "123"}

	response, err := svc.DeleteChallenge(c, req)
	assert.Error(t, err)
	assert.Nil(t, response)
}
