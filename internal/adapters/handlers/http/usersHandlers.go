package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	model "CrudPlatform/internal/core/domain/repository/model/users"

	"CrudPlatform/internal/core/ports"
)

type managementHandler struct {
	Service    ports.CommunicationUserServices
	Repository ports.DBRepository
}

func newHandler(service ports.CommunicationUserServices, repo ports.DBRepository) *managementHandler {
	return &managementHandler{
		Service:    service,
		Repository: repo,
	}
}

func (o *managementHandler) postUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var User model.User
		if err := c.BindJSON(&User); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Request"})
			return
		}
		entityResponse, err := o.Service.CreateUser(c, &User)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.Set("entityResponse", *entityResponse)
		c.JSON(http.StatusOK, entityResponse)
	}
}

func (o *managementHandler) getUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var User model.GetUser
		User.Id = c.Param("id")
		if err := c.ShouldBindQuery(&User); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Request "})
			return
		}
		entityResponse, err := o.Service.SelectUser(c, &User)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.Set("entityResponse", *entityResponse)
		c.JSON(http.StatusOK, entityResponse)
	}
}

func (o *managementHandler) putUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var User model.UpdateUser
		User.Id = c.Param("id")
		if err := c.ShouldBindQuery(&User); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Request "})
			return
		}
		if err := c.BindJSON(&User); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Request "})
			return
		}
		entityResponse, err := o.Service.UpdateUser(c, &User)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.Set("entityResponse", *entityResponse)
		c.JSON(http.StatusOK, entityResponse)
	}
}

func (o *managementHandler) deleteUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var User model.DeleteUser
		User.Id = c.Param("id")
		if err := c.ShouldBindQuery(&User); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Request "})
			return
		}
		entityResponse, err := o.Service.DeleteUser(c, &User)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.Set("entityResponse", *entityResponse)
		c.JSON(http.StatusOK, entityResponse)
	}
}
