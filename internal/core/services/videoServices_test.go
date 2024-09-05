package service

import (
	"errors"
	"net/http/httptest"
	"testing"

	entity "CrudPlatform/internal/core/domain/repository"
	model "CrudPlatform/internal/core/domain/repository/model/videos"
	schema "CrudPlatform/internal/core/domain/repository/schema/videos"
	mockRepository "CrudPlatform/internal/core/ports/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewServiceVideo(t *testing.T) {
	mockRepo := mockRepository.NewDBRepositoryVideo(t)
	service := NewServiceVideo(mockRepo)
	assert.NotNil(t, service, "El servicio no debe ser nil")
}

func TestCreateVideo(t *testing.T) {
	mockRepo := mockRepository.NewDBRepositoryVideo(t)
	svc := &RepositoryVideo{repo: mockRepo}

	mockRepo.On("CreateVideo", mock.Anything, mock.Anything).Return("123", nil)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := &model.Videos{
		Title:       "Test Video",
		Description: "Test Description",
	}

	expectedResp := &entity.Response{
		Data: "123",
		Result: entity.Result{
			Details: []entity.Detail{
				{InternalCode: "200", Message: "OK", Detail: "Registro Creado"},
			},
			Source: "Create Video",
		},
	}

	response, err := svc.CreateVideo(c, req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp, response)
}

func TestCreateVideo_ErrorCase(t *testing.T) {
	mockRepo := mockRepository.NewDBRepositoryVideo(t)
	svc := &RepositoryVideo{repo: mockRepo}

	mockRepo.On("CreateVideo", mock.Anything, mock.Anything).Return("", errors.New("error simulado"))

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := &model.Videos{
		Title:       "Test Video",
		Description: "Test Description",
	}

	response, err := svc.CreateVideo(c, req)
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestSelectVideo(t *testing.T) {
	mockRepo := mockRepository.NewDBRepositoryVideo(t)
	svc := &RepositoryVideo{repo: mockRepo}

	mockResponse := &schema.VideosGetResponse{
		Title:       "Test Video",
		Description: "Test Description",
	}
	mockRepo.On("SelectVideo", mock.Anything, mock.Anything).Return(mockResponse, nil)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := &model.GetVideo{ID: "123"}

	expectedResp := &entity.Response{
		Data: mockResponse,
		Result: entity.Result{
			Details: []entity.Detail{
				{InternalCode: "200", Message: "OK", Detail: "Registro Seleccionado"},
			},
			Source: "Select Video",
		},
	}

	response, err := svc.SelectVideo(c, req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp, response)
}

func TestSelectVideo_ErrorCase(t *testing.T) {
	mockRepo := mockRepository.NewDBRepositoryVideo(t)
	svc := &RepositoryVideo{repo: mockRepo}

	mockRepo.On("SelectVideo", mock.Anything, mock.Anything).Return(nil, errors.New("error simulado"))

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := &model.GetVideo{ID: "123"}

	response, err := svc.SelectVideo(c, req)
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestUpdateVideo(t *testing.T) {
	mockRepo := mockRepository.NewDBRepositoryVideo(t)
	svc := &RepositoryVideo{repo: mockRepo}

	mockResponse := &schema.VideosUpdateResponse{
		Title:       "Updated Video",
		Description: "Updated Description",
	}
	mockRepo.On("UpdateVideo", mock.Anything, mock.Anything).Return(mockResponse, nil)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := &model.UpdateVideo{
		ID:          "123",
		Title:       "Updated Video",
		Description: "Updated Description",
	}

	expectedResp := &entity.Response{
		Data: mockResponse,
		Result: entity.Result{
			Details: []entity.Detail{
				{InternalCode: "200", Message: "OK", Detail: "Registro Actualizado"},
			},
			Source: "Update Video",
		},
	}

	response, err := svc.UpdateVideo(c, req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp, response)
}

func TestUpdateVideo_ErrorCase(t *testing.T) {
	mockRepo := mockRepository.NewDBRepositoryVideo(t)
	svc := &RepositoryVideo{repo: mockRepo}

	mockRepo.On("UpdateVideo", mock.Anything, mock.Anything).Return(nil, errors.New("error simulado"))

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := &model.UpdateVideo{
		ID:          "123",
		Title:       "Updated Video",
		Description: "Updated Description",
	}

	response, err := svc.UpdateVideo(c, req)
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestDeleteVideo(t *testing.T) {
	mockRepo := mockRepository.NewDBRepositoryVideo(t)
	svc := &RepositoryVideo{repo: mockRepo}

	mockRepo.On("DeleteVideo", mock.Anything, mock.Anything).Return(nil)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := &model.DeleteVideo{ID: "123"}

	expectedResp := &entity.Response{
		Result: entity.Result{
			Details: []entity.Detail{
				{InternalCode: "200", Message: "OK", Detail: "Registro Eliminado"},
			},
			Source: "Delete Challenge",
		},
	}

	response, err := svc.DeleteVideo(c, req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp, response)
}

func TestDeleteVideo_ErrorCase(t *testing.T) {
	mockRepo := mockRepository.NewDBRepositoryVideo(t)
	svc := &RepositoryVideo{repo: mockRepo}

	mockRepo.On("DeleteVideo", mock.Anything, mock.Anything).Return(errors.New("error simulado"))

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := &model.DeleteVideo{ID: "123"}

	response, err := svc.DeleteVideo(c, req)
	assert.Error(t, err)
	assert.Nil(t, response)
}
