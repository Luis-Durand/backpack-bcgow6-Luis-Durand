package main

import (
	"github.com/Luis-Durand/backpack-bcgow6-Luis-Durand/go-testing-final/cmd/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.MapRoutes(r)

	r.Run(":18085")

}
