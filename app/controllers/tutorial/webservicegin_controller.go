package tutorial

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// album - data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums - album data slice
var albums = []album{
	{ID: "1", Title: "Jay", Artist: "Jay Chou", Price: 35.5},
	{ID: "2", Title: "小满", Artist: "音阙诗听", Price: 2.5},
	{ID: "3", Title: "离人赋", Artist: "云汐", Price: 5.0},
	{ID: "4", Title: "从前说", Artist: "小阿七", Price: 1.8},
}

// getAlbums - responds albums slice
func getAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, albums)
}

// postAlbums - add a new album
func postAlbums(context *gin.Context) {
	var newAlbum album
	if err := context.BindJSON(&newAlbum); err != nil {
		return
	}
	// check album if exist
	for _, al := range albums {
		if al.ID == newAlbum.ID {
			context.IndentedJSON(http.StatusForbidden, gin.H{
				"message": fmt.Sprintf("album: %s already exists", al.ID),
			})
			return
		}
	}
	albums = append(albums, newAlbum)
	context.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID - get the album with specific id
func getAlbumByID(context *gin.Context) {
	id := context.Query("id")
	fmt.Println("id", id)
	for _, al := range albums {
		if al.ID == id {
			context.IndentedJSON(http.StatusOK, al)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{
		"message": fmt.Sprintf("not found album with id: %s", id),
	})
}
