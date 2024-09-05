package http

import (
	"CrudPlatform/internal/adapters/handlers/http/middleware"
	"time"

	"database/sql"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func CreateServer(db *sql.DB) *gin.Engine {
	server := gin.Default()

	server.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET,POST,DELETE,PUT",
		RequestHeaders: "Origin, Authorization, Content-Type, Access-Control-Allow-Origin",
		MaxAge:         50 * time.Second,
	}))

	server.Use(middleware.AuthenticationMiddleware())

	RegisterRoutes(server, db)

	return server
}

func RunServer(db *sql.DB) {
	server := CreateServer(db)
	server.Run(":8086")
}
