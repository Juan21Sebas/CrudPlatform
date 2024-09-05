package http

import (
	repository "CrudPlatform/internal/adapters/repository"
	services "CrudPlatform/internal/core/services"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(e *gin.Engine, db *sql.DB) {
	// Crea e inicializa el repositorio BDRepository con la conexión a la base de datos
	Repository := repository.NewBdRepository(db)
	RepositoryChallenge := repository.NewBdRepositoryChallenge(db)
	RepositoryVideo := repository.NewBdRepositoryVideo(db)

	// Crea e inicializa el servicio con el repositorio
	Service := services.NewService(Repository)
	ServiceChallenge := services.NewServiceChallenge(RepositoryChallenge)
	ServiceVideo := services.NewServiceVideo(RepositoryVideo)

	// Crea el manejador con el servicio y el repositorio
	managementHandler := newHandler(Service, Repository)
	managementChallengeHandler := newChallengeHandler(ServiceChallenge, RepositoryChallenge)
	managementVideoHandler := newVideosHandler(ServiceVideo, RepositoryVideo)

	// Registra las rutas Users
	e.POST("/users/", managementHandler.postUsers())
	e.GET("/users/:id", managementHandler.getUsers())
	e.PUT("/users/:id", managementHandler.putUsers())
	e.DELETE("/users/:id", managementHandler.deleteUsers())

	// Registra las rutas Challenge
	e.POST("/challenge/", managementChallengeHandler.postChallenge())
	e.GET("/challenge/:id", managementChallengeHandler.getChallenge())
	e.PUT("/challenge/:id", managementChallengeHandler.putChallenge())
	e.DELETE("/challenge/:id", managementChallengeHandler.deleteChallenge())

	// Registra las rutas Video
	e.POST("/video/", managementVideoHandler.postVideo())
	e.GET("/video/:id", managementVideoHandler.getVideo())
	e.PUT("/video/:id", managementVideoHandler.putVideo())
	e.DELETE("/video/:id", managementVideoHandler.deleteVideo())

}
