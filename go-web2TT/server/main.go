package main

import (
	"github.com/Luis-Durand/backpack-bcgow6-Luis-Durand/go-web2TT/internal/products"
	"github.com/Luis-Durand/backpack-bcgow6-Luis-Durand/go-web2TT/pkg/store"
	"github.com/Luis-Durand/backpack-bcgow6-Luis-Durand/go-web2TT/server/handler"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		err.Error()
		return
	}
	db := store.NewStore(store.FileType, "./products.json")
	repo := products.NewRepository(db)
	data := products.NewService(repo)
	serv := handler.NewHandler(data)

	router := gin.Default()

	pr := router.Group("/products")

	pr.POST("/", serv.Store())
	pr.GET("/", serv.GetAll())
	pr.PUT("/:id", serv.Update())
	pr.DELETE("/:id", serv.Delete())
	pr.PATCH("/:id", serv.UpdateNameAndPrice())
	router.Run()

}
