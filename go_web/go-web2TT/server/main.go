package main

import (
	"os"

	"github.com/Luis-Durand/backpack-bcgow6-Luis-Durand/go_web/go-web2TT/docs"
	"github.com/Luis-Durand/backpack-bcgow6-Luis-Durand/go_web/go-web2TT/internal/products"
	"github.com/Luis-Durand/backpack-bcgow6-Luis-Durand/go_web/go-web2TT/pkg/store"
	"github.com/Luis-Durand/backpack-bcgow6-Luis-Durand/go_web/go-web2TT/server/handler"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title          Bootcamp Go Wave 6 - API
// @version        1.0
// @description    This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name  API Practica
// @contact.url   http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host     localhost:8080
// @BasePath /api/v1
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
	api := router.Group("/api/v1")

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	pr := api.Group("/products")

	pr.POST("/", serv.Store())
	pr.GET("/", serv.GetAll())
	pr.PUT("/:id", serv.Update())
	pr.DELETE("/:id", serv.Delete())
	pr.PATCH("/:id", serv.UpdateNameAndPrice())
	router.Run()

}
