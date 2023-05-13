package helloworld

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func runGet(path string, handler gin.HandlerFunc) {
	router := gin.Default()
	router.GET(path, handler)
	router.Run(":9090")
}

// TestHelloWorld - HelloWorld controller testing
func TestHelloWorld(t *testing.T) {
	runGet("/hw", HelloWorld)
}
