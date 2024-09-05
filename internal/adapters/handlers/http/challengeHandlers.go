package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	model "CrudPlatform/internal/core/domain/repository/model/challenges"
	"CrudPlatform/internal/core/ports"
)

type managementChallengeHandler struct {
	Service    ports.CommunicationChallengeServices
	Repository ports.DBRepositoryChallenge
}

func newChallengeHandler(service ports.CommunicationChallengeServices, repo ports.DBRepositoryChallenge) *managementChallengeHandler {
	return &managementChallengeHandler{
		Service:    service,
		Repository: repo,
	}
}

func (o *managementChallengeHandler) postChallenge() gin.HandlerFunc {
	return func(c *gin.Context) {
		var User model.Challenge
		if err := c.BindJSON(&User); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Request"})
			return
		}
		entityResponse, err := o.Service.CreateChallenge(c, &User)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.Set("entityResponse", *entityResponse)
		c.JSON(http.StatusOK, entityResponse)
	}
}

func (o *managementChallengeHandler) getChallenge() gin.HandlerFunc {
	return func(c *gin.Context) {
		var User model.GetChallenge
		User.ID = c.Param("id")
		if err := c.ShouldBindQuery(&User); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Request "})
			return
		}
		entityResponse, err := o.Service.SelectChallenge(c, &User)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.Set("entityResponse", *entityResponse)
		c.JSON(http.StatusOK, entityResponse)
	}
}

func (o *managementChallengeHandler) putChallenge() gin.HandlerFunc {
	return func(c *gin.Context) {
		var User model.UpdateChallenge
		User.ID = c.Param("id")
		if err := c.ShouldBindQuery(&User); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Request "})
			return
		}
		if err := c.BindJSON(&User); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Request "})
			return
		}
		entityResponse, err := o.Service.UpdateChallenge(c, &User)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.Set("entityResponse", *entityResponse)
		c.JSON(http.StatusOK, entityResponse)
	}
}

func (o *managementChallengeHandler) deleteChallenge() gin.HandlerFunc {
	return func(c *gin.Context) {
		var User model.DeleteChallenge
		User.ID = c.Param("id")
		if err := c.ShouldBindQuery(&User); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Request "})
			return
		}
		entityResponse, err := o.Service.DeleteChallenge(c, &User)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.Set("entityResponse", *entityResponse)
		c.JSON(http.StatusOK, entityResponse)
	}
}
