package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"testing"
)

type api struct {
	method  string
	path    string
	handler gin.HandlerFunc
}

var apis = []api{
	{method: http.MethodGet, path: "/albums", handler: getAlbums},
	{method: http.MethodGet, path: "/album", handler: getAlbumByID},
	{method: http.MethodPost, path: "/addAlbum", handler: postAlbums},
}

func runApis(apis []api) {
	router := gin.Default()
	for _, api := range apis {
		router.Handle(api.method, api.path, api.handler)
	}
	err := router.Run(":9090")
	if err != nil {
		log.Fatalln(err)
	}
}

// getAlbums testing
func TestAlbums(t *testing.T) {
	runApis(apis)
}
