package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stepan41k/GinTest/pkg/endpoints"
)


func main() {
	router := gin.Default()

	router.GET("/albums", endpoints.GetAlbums)
	router.GET("/albums/:id", endpoints.GetAlbumByID)
	router.POST("/albums", endpoints.PostAlbum)
	router.POST("/albums/:id", endpoints.UpdateAlbum)
	router.DELETE("/albums/:id", endpoints.DeleteAlbum)

	router.Run("localhost:8080")
	
}



