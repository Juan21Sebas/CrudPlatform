package service

import (
	"CrudPlatform/internal/core/ports"
	"net/http"
	"strconv"

	entity "CrudPlatform/internal/core/domain/repository"
	model "CrudPlatform/internal/core/domain/repository/model/challenges"

	"github.com/gin-gonic/gin"
)

type RepositoryChallenge struct {
	repo ports.DBRepositoryChallenge
}

func NewServiceChallenge(repo ports.DBRepositoryChallenge) *RepositoryChallenge {
	return &RepositoryChallenge{
		repo: repo,
	}
}

func (r *RepositoryChallenge) CreateChallenge(ctx *gin.Context, request *model.Challenge) (*entity.Response, error) {

	resp, err := r.repo.CreateChallenge(ctx, request)
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
			Source: "Create Challenge",
		},
	}, nil

}

func (r *RepositoryChallenge) SelectChallenge(ctx *gin.Context, request *model.GetChallenge) (*entity.Response, error) {

	resp, err := r.repo.SelectChallenge(ctx, request)
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
			Source: "Select Challenge",
		},
	}, nil

}

func (r *RepositoryChallenge) UpdateChallenge(ctx *gin.Context, request *model.UpdateChallenge) (*entity.Response, error) {

	resp, err := r.repo.UpdateChallenge(ctx, request)
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
			Source: "Update Challenge",
		},
	}, nil

}

func (r *RepositoryChallenge) DeleteChallenge(ctx *gin.Context, request *model.DeleteChallenge) (*entity.Response, error) {

	err := r.repo.DeleteChallenge(ctx, request)
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
