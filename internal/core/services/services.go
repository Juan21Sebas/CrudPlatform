package service

import (
	"CrudPlatform/internal/core/ports"
	"net/http"
	"strconv"

	entity "CrudPlatform/internal/core/domain/repository"
	model "CrudPlatform/internal/core/domain/repository/model/users"

	"github.com/gin-gonic/gin"
)

type Repository struct {
	repo ports.DBRepository
}

func NewService(repo ports.DBRepository) *Repository {
	return &Repository{
		repo: repo,
	}
}

func (r *Repository) CreateUser(ctx *gin.Context, request *model.User) (*entity.Response, error) {

	resp, err := r.repo.CreateUser(ctx, request)
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
			Source: "Create User",
		},
	}, nil

}

func (r *Repository) SelectUser(ctx *gin.Context, request *model.GetUser) (*entity.Response, error) {

	resp, err := r.repo.SelectUser(ctx, request)
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
			Source: "Select User",
		},
	}, nil

}

func (r *Repository) UpdateUser(ctx *gin.Context, request *model.UpdateUser) (*entity.Response, error) {

	resp, err := r.repo.UpdateUser(ctx, request)
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
			Source: "Update User",
		},
	}, nil

}

func (r *Repository) DeleteUser(ctx *gin.Context, request *model.DeleteUser) (*entity.Response, error) {

	err := r.repo.DeleteUser(ctx, request)
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
			Source: "Delete User",
		},
	}, nil

}
