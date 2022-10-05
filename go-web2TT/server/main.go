package main

import (
	"github.com/Luis-Durand/backpack-bcgow6-Luis-Durand/go-web2TT/internal/products"
	"github.com/Luis-Durand/backpack-bcgow6-Luis-Durand/go-web2TT/server/handler"
	"github.com/gin-gonic/gin"
)

func main() {

	repo := products.NewRepository()
	data := products.NewService(repo)
	serv := handler.NewHandler(data)
	router := gin.Default()

	pr := router.Group("/products")

	pr.POST("/", serv.Store())
	pr.GET("/", serv.GetAll())
	pr.PUT("/:id", serv.Update())

	router.Run()

}
