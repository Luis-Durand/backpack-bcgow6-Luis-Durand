package handler

import (
	"net/http"
	"strconv"

	"github.com/Luis-Durand/backpack-bcgow6-Luis-Durand/go-web2TT/internal/products"
	"github.com/gin-gonic/gin"
)

type request struct {
	Name   string `json:"name" binding:"required"`
	Type   string `json:"type" binding:"required"`
	Amount int    `json:"amount" binding:"required"`
	Price  int    `json:"price" binding:"required"`
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
		if token != "bootcamp" {
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
		if token != "bootcamp" {
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
		if token != "bootcamp" {
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

		/* 	data,err :=ctx.Param("id") */
		numId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		prods, err := s.service.Update(int(numId), req.Name, req.Type, req.Amount, req.Price)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, prods)
	}
}
