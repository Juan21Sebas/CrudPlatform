package ports

import (
	entity "CrudPlatform/internal/core/domain/repository"
	model "CrudPlatform/internal/core/domain/repository/model/users"

	schema "CrudPlatform/internal/core/domain/repository/schema/users"

	"github.com/gin-gonic/gin"
)

type CommunicationUserServices interface {
	CreateUser(ctx *gin.Context, request *model.User) (*entity.Response, error)
	SelectUser(ctx *gin.Context, request *model.GetUser) (*entity.Response, error)
	UpdateUser(ctx *gin.Context, request *model.UpdateUser) (*entity.Response, error)
	DeleteUser(ctx *gin.Context, request *model.DeleteUser) (*entity.Response, error)
}

type DBRepository interface {
	CreateUser(ctx *gin.Context, request *model.User) (string, error)
	SelectUser(ctx *gin.Context, request *model.GetUser) (*schema.UsersGetResponse, error)
	UpdateUser(ctx *gin.Context, request *model.UpdateUser) (*schema.UsersUpdateResponse, error)
	DeleteUser(ctx *gin.Context, request *model.DeleteUser) error
}
