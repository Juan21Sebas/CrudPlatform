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

	// Crea e inicializa el servicio con el repositorio
	Service := services.NewService(Repository)

	// Crea el manejador con el servicio y el repositorio
	managementHandler := newHandler(Service, Repository)

	// Registra las rutas Users
	e.POST("/users/", managementHandler.postUsers())
	e.GET("/users/:id", managementHandler.getUsers())
	e.PUT("/users/:id", managementHandler.putUsers())
	e.DELETE("/users/:id", managementHandler.deleteUsers())
}
