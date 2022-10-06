package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/Luis-Durand/backpack-bcgow6-Luis-Durand/go-web2TT/internal/products"
	"github.com/gin-gonic/gin"
)

type request struct {
	Name   string `json:"name" `
	Type   string `json:"type" `
	Amount int    `json:"amount" `
	Price  int    `json:"price" `
}

type Product struct {
	service products.Service
}

func NewHandler(p products.Service) *Product {
	return &Product{
		service: p,
	}

}

func (s *Product) GetAll() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusBadRequest, gin.H{

				"error": "Token invalido",
			})
			return
		}

		prods, err := s.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, prods)

	}

}

func (s *Product) Store() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusBadRequest, gin.H{

				"error": "Token invalido",
			})
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if req.Name == "" {

			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "El nombre del producto es necesario",
			})
			return
		}
		if req.Type == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "El tipo del producto es necesario",
			})
			return

		}
		if req.Amount == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "La cantidad del producto es necesario",
			})
			return

		}
		if req.Price == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "El precio del producto es necesario",
			})
			return

		}

		prod, err := s.service.Store(req.Name, req.Type, req.Amount, req.Price)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, prod)
	}

}

func (s *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Token invalido",
			})
			return
		}
		var req request
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if req.Name == "" {

			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "El nombre del producto es necesario",
			})
			return
		}
		if req.Type == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "El tipo del producto es necesario",
			})
			return

		}
		if req.Amount == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "La cantidad del producto es necesario",
			})
			return

		}
		if req.Price == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "El precio del producto es necesario",
			})
			return

		}
		numId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		prods, err := s.service.Update(int(numId), req.Name, req.Type, req.Amount, req.Price)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, prods)
	}
}

func (s *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token invalido",
			})
			return
		}

		idParam := ctx.Param("id")

		paramInt, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		newErr := s.service.Delete(paramInt)

		if newErr != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": newErr.Error(),
			})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{
			"data": fmt.Sprintf("El producto %d a sido eliminado correctamente", paramInt),
		})
	}
}

func (s *Product) UpdateNameAndPrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token invalido",
			})
			return
		}

		idParam := ctx.Param("id")
		paramInt, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var req request
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if req.Name == "" {

			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "El nombre del producto es necesario",
			})
			return
		}

		if req.Price == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "El precio del producto es necesario",
			})
			return

		}

		prod, err := s.service.UpdateNameAndPrice(paramInt, req.Name, req.Price)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, prod)
	}
}
