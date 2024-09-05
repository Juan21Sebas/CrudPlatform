package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	model "CrudPlatform/internal/core/domain/repository/model/videos"
	"CrudPlatform/internal/core/ports"
)

type managementVideoHandler struct {
	Service    ports.CommunicationVideoServices
	Repository ports.DBRepositoryVideo
}

func newVideosHandler(service ports.CommunicationVideoServices, repo ports.DBRepositoryVideo) *managementVideoHandler {
	return &managementVideoHandler{
		Service:    service,
		Repository: repo,
	}
}

func (o *managementVideoHandler) postVideo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var User model.Videos
		if err := c.BindJSON(&User); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Request"})
			return
		}
		entityResponse, err := o.Service.CreateVideo(c, &User)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.Set("entityResponse", *entityResponse)
		c.JSON(http.StatusOK, entityResponse)
	}
}

func (o *managementVideoHandler) getVideo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var User model.GetVideo
		User.ID = c.Param("id")
		if err := c.ShouldBindQuery(&User); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Request "})
			return
		}
		entityResponse, err := o.Service.SelectVideo(c, &User)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.Set("entityResponse", *entityResponse)
		c.JSON(http.StatusOK, entityResponse)
	}
}

func (o *managementVideoHandler) putVideo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var User model.UpdateVideo
		User.ID = c.Param("id")
		if err := c.ShouldBindQuery(&User); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Request "})
			return
		}
		if err := c.BindJSON(&User); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Request "})
			return
		}
		entityResponse, err := o.Service.UpdateVideo(c, &User)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.Set("entityResponse", *entityResponse)
		c.JSON(http.StatusOK, entityResponse)
	}
}

func (o *managementVideoHandler) deleteVideo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var User model.DeleteVideo
		User.ID = c.Param("id")
		if err := c.ShouldBindQuery(&User); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Request "})
			return
		}
		entityResponse, err := o.Service.DeleteVideo(c, &User)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.Set("entityResponse", *entityResponse)
		c.JSON(http.StatusOK, entityResponse)
	}
}
