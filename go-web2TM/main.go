package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Products struct {
	Id     int    `json:"id"`
	Name   string `json:"name"  `
	Type   string `json:"type" `
	Price  int    `json:"price" `
	Amount int    `json:"amount" `
}

var arrProducts []Products

func main() {
	router := gin.Default()

	pr := router.Group("/products")
	pr.POST("/", Save)

	router.Run()

}

func Save(c *gin.Context) {

	token := c.GetHeader("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No se ha enviado token",
		})
		return
	}
	var prod Products

	if token != "bootcamp" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "“no tiene permisos para realizar la petición solicitada”.",
		})
		return
	}

	if err := c.ShouldBindJSON(&prod); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	validate := validator.New()
	/* err:=validator.FieldError.Tag(prod) */
	err := validate.Struct(prod)

	if err != nil {

		for _, err := range err.(validator.ValidationErrors) {

			c.String(http.StatusBadRequest, "el campo %s es requerido \n", err.Field())

		}

		// from here you can create your own error messages in whatever language you wish

		return
	}

	prod.Id = len(arrProducts) + 1
	arrProducts = append(arrProducts, prod)

	c.JSON(http.StatusOK, prod)

}
