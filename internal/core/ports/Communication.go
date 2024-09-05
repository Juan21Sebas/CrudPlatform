package ports

import (
	entity "CrudPlatform/internal/core/domain/repository"
	modelChallenge "CrudPlatform/internal/core/domain/repository/model/challenges"
	model "CrudPlatform/internal/core/domain/repository/model/users"
	modelVideo "CrudPlatform/internal/core/domain/repository/model/videos"

	schemaChallenges "CrudPlatform/internal/core/domain/repository/schema/challenges"
	schema "CrudPlatform/internal/core/domain/repository/schema/users"
	schemaVideos "CrudPlatform/internal/core/domain/repository/schema/videos"

	"github.com/gin-gonic/gin"
)

type CommunicationUserServices interface {
	CreateUser(ctx *gin.Context, request *model.User) (*entity.Response, error)
	SelectUser(ctx *gin.Context, request *model.GetUser) (*entity.Response, error)
	UpdateUser(ctx *gin.Context, request *model.UpdateUser) (*entity.Response, error)
	DeleteUser(ctx *gin.Context, request *model.DeleteUser) (*entity.Response, error)
}

type CommunicationChallengeServices interface {
	CreateChallenge(ctx *gin.Context, request *modelChallenge.Challenge) (*entity.Response, error)
	SelectChallenge(ctx *gin.Context, request *modelChallenge.GetChallenge) (*entity.Response, error)
	UpdateChallenge(ctx *gin.Context, request *modelChallenge.UpdateChallenge) (*entity.Response, error)
	DeleteChallenge(ctx *gin.Context, request *modelChallenge.DeleteChallenge) (*entity.Response, error)
}

type CommunicationVideoServices interface {
	CreateVideo(ctx *gin.Context, request *modelVideo.Videos) (*entity.Response, error)
	SelectVideo(ctx *gin.Context, request *modelVideo.GetVideo) (*entity.Response, error)
	UpdateVideo(ctx *gin.Context, request *modelVideo.UpdateVideo) (*entity.Response, error)
	DeleteVideo(ctx *gin.Context, request *modelVideo.DeleteVideo) (*entity.Response, error)
}

type DBRepositoryUsers interface {
	CreateUser(ctx *gin.Context, request *model.User) (string, error)
	SelectUser(ctx *gin.Context, request *model.GetUser) (*schema.UsersGetResponse, error)
	UpdateUser(ctx *gin.Context, request *model.UpdateUser) (*schema.UsersUpdateResponse, error)
	DeleteUser(ctx *gin.Context, request *model.DeleteUser) error
}

type DBRepositoryChallenge interface {
	CreateChallenge(ctx *gin.Context, request *modelChallenge.Challenge) (string, error)
	SelectChallenge(ctx *gin.Context, request *modelChallenge.GetChallenge) (*schemaChallenges.ChallengeGetResponse, error)
	UpdateChallenge(ctx *gin.Context, request *modelChallenge.UpdateChallenge) (*schemaChallenges.ChallengeUpdateResponse, error)
	DeleteChallenge(ctx *gin.Context, request *modelChallenge.DeleteChallenge) error
}

type DBRepositoryVideo interface {
	CreateVideo(ctx *gin.Context, request *modelVideo.Videos) (string, error)
	SelectVideo(ctx *gin.Context, request *modelVideo.GetVideo) (*schemaVideos.VideosGetResponse, error)
	UpdateVideo(ctx *gin.Context, request *modelVideo.UpdateVideo) (*schemaVideos.VideosUpdateResponse, error)
	DeleteVideo(ctx *gin.Context, request *modelVideo.DeleteVideo) error
}
