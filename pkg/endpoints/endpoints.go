package endpoints

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/stepan41k/GinTest/pkg/models"

)

func GetAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, models.Albums)
}

func GetAlbumByID(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, a := range models.Albums {
		if a.ID == id {
			ctx.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func PostAlbum(ctx *gin.Context) {
	var newAlbum models.Album

	if err := ctx.ShouldBindJSON(&newAlbum); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.Albums = append(models.Albums, newAlbum)
	ctx.JSON(http.StatusCreated, gin.H{"message" : "album added"})
}

func UpdateAlbum(ctx *gin.Context) {
	id := ctx.Param("id")

	var updateAlbum models.Album

	if err := ctx.ShouldBindJSON(&updateAlbum); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, album := range models.Albums {
		if album.ID == id {
			if updateAlbum.Artist != "" {
				models.Albums[i].Artist = updateAlbum.Artist
			}
			if updateAlbum.Price != 0 {
				models.Albums[i].Price = updateAlbum.Price
			}
			if updateAlbum.Title != "" {
				models.Albums[i].Title = updateAlbum.Title
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "album updated"})
			return
			
		}
	}
}

func DeleteAlbum(ctx *gin.Context) {
	id := ctx.Param("id")

	for i, album := range models.Albums {
		if album.ID == id {
			models.Albums = append(models.Albums[:i], models.Albums[i+1:]...)
			ctx.JSON(http.StatusOK, gin.H{"message": "album removed"})
			return
		}
	}
}