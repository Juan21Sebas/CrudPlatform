package service

import (
	"CrudPlatform/internal/core/ports"
	"net/http"
	"strconv"

	entity "CrudPlatform/internal/core/domain/repository"
	model "CrudPlatform/internal/core/domain/repository/model/videos"

	"github.com/gin-gonic/gin"
)

type RepositoryVideo struct {
	repo ports.DBRepositoryVideo
}

func NewServiceVideo(repo ports.DBRepositoryVideo) *RepositoryVideo {
	return &RepositoryVideo{
		repo: repo,
	}
}

func (r *RepositoryVideo) CreateVideo(ctx *gin.Context, request *model.Videos) (*entity.Response, error) {

	resp, err := r.repo.CreateVideo(ctx, request)
	if err != nil {
		return nil, err
	}

	return &entity.Response{
		Data: resp,
		Result: entity.Result{
			Details: []entity.Detail{
				{
					InternalCode: strconv.Itoa(http.StatusOK),
					Message:      http.StatusText(http.StatusOK),
					Detail:       "Registro Creado",
				},
			},
			Source: "Create Video",
		},
	}, nil

}

func (r *RepositoryVideo) SelectVideo(ctx *gin.Context, request *model.GetVideo) (*entity.Response, error) {

	resp, err := r.repo.SelectVideo(ctx, request)
	if err != nil {
		return nil, err
	}

	return &entity.Response{
		Data: resp,
		Result: entity.Result{
			Details: []entity.Detail{
				{
					InternalCode: strconv.Itoa(http.StatusOK),
					Message:      http.StatusText(http.StatusOK),
					Detail:       "Registro Seleccionado",
				},
			},
			Source: "Select Video",
		},
	}, nil

}

func (r *RepositoryVideo) UpdateVideo(ctx *gin.Context, request *model.UpdateVideo) (*entity.Response, error) {

	resp, err := r.repo.UpdateVideo(ctx, request)
	if err != nil {
		return nil, err
	}

	return &entity.Response{
		Data: resp,
		Result: entity.Result{
			Details: []entity.Detail{
				{
					InternalCode: strconv.Itoa(http.StatusOK),
					Message:      http.StatusText(http.StatusOK),
					Detail:       "Registro Actualizado",
				},
			},
			Source: "Update Video",
		},
	}, nil

}

func (r *RepositoryVideo) DeleteVideo(ctx *gin.Context, request *model.DeleteVideo) (*entity.Response, error) {

	err := r.repo.DeleteVideo(ctx, request)
	if err != nil {
		return nil, err
	}

	return &entity.Response{
		Result: entity.Result{
			Details: []entity.Detail{
				{
					InternalCode: strconv.Itoa(http.StatusOK),
					Message:      http.StatusText(http.StatusOK),
					Detail:       "Registro Eliminado",
				},
			},
			Source: "Delete Challenge",
		},
	}, nil

}
