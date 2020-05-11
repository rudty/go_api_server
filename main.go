package main

import (
	"api/album"
	"api/common/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.Recovery)

	g := r.Group("/api")
	album.Register(g)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
