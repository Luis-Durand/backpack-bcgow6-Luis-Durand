package handler

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Luis-Durand/backpack-bcgow6-Luis-Durand/go-test/test2-TM/internal/products"
	"github.com/Luis-Durand/backpack-bcgow6-Luis-Durand/go-test/test2-TM/pkg/web"
	"github.com/gin-gonic/gin"
)

type Request struct {
	Name   string `json:"name" `
	Type   string `json:"type" `
	Price  int    `json:"price" `
	Amount int    `json:"amount" `
}

type Product struct {
	service products.Service
}

func NewHandler(p products.Service) *Product {
	return &Product{
		service: p,
	}

}

// GetAll godoc
// @Summary     Show list products
// @Tags     Products
// @Description get products
// @Produce     json
// @Param       token header   string       true "TOKEN"
// @Success     200   {object} web.Response "List Products"
// @Failure     401   {object} web.Response "Unauthorized"
// @Failure     500   {object} web.Response "Internal server error"
// @Failure     404   {object} web.Response "Not found products"
// @Router      /products [GET]
func (s *Product) GetAll() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "Token invalido"))
			return
		}

		prods, err := s.service.GetAll()
		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
			return
		}

		/* if len(prods) == 0 {
			ctx.JSON(404, web.NewResponse(404, nil, "No  hay productos almacenados"))
			return
		} */

		ctx.JSON(200, web.NewResponse(200, prods, ""))

	}

}

// Store godoc
// @Summary     Post product
// @Tags     Products
// @Description post product
// @Accept      json
// @Produce     json
// @Param       token   header   string       true "TOKEN"
// @Param       product body     Request      true "Product to update"
// @Success     200     {object} web.Response "List Products"
// @Failure     400     {object} web.Response "Bad Request"
// @Failure     401     {object} web.Response "Unauthorized"
// @Router      /products [POST]
func (s *Product) Store() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "Token invalido"))
			return
		}
		var req Request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		if req.Name == "" {

			ctx.JSON(400, web.NewResponse(400, nil, "El nombre del producto es necesario"))
			return
		}
		if req.Type == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El tipo del producto es necesario"))
			return

		}
		if req.Amount == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "La cantidad del producto es necesario"))
			return

		}
		if req.Price == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "El precio del producto es necesario"))
			return

		}

		prod, err := s.service.Store(req.Name, req.Type, req.Amount, req.Price)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, prod, ""))
	}

}

// Update godoc
// @Summary     Update product products
// @Tags     Products
// @Description Update product
// @Accept      json
// @Produce     json
// @Param       id      path     int          true "Id product"
// @Param       token   header   string       true "TOKEN"
// @Param       product body     Request      true "Product to update"
// @Success     200     {object} web.Response "List Products"
// @Failure     400     {object} web.Response "Bad Request"
// @Failure     401     {object} web.Response "Unauthorized"
// @Router      /products/{id} [PUT]
func (s *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "Token invalido"))
			return
		}
		var req Request
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		if req.Name == "" {

			ctx.JSON(400, web.NewResponse(400, nil, "El nombre del producto es necesario"))
			return
		}
		if req.Type == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El tipo del produto es necesario"))
			return

		}
		if req.Amount == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "La cantidad del producto es necesario"))
			return

		}
		if req.Price == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "El precio del producto es necesario"))
			return

		}
		numId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		prods, err := s.service.Update(int(numId), req.Name, req.Type, req.Price, req.Amount)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, prods, ""))
	}
}

// Delete godoc
// @Summary     Delete product
// @Tags     Products
// @Description delete product
// @Produce     json
// @Param       token header   string       true "TOKEN"
// @Param       id    path     int          true "Id product"
// @Success     200   {object} web.Response "List Products"
// @Failure     400   {object} web.Response "Bad Request"
// @Failure     401   {object} web.Response "Unauthorized"
// @Router      /products/{id} [DELETE]
func (s *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "Token invalido"))
			return
		}

		idParam := ctx.Param("id")

		paramInt, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		newErr := s.service.Delete(paramInt)

		if newErr != nil {
			ctx.JSON(400, web.NewResponse(400, nil, newErr.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, nil, fmt.Sprintf("El producto %d a sido eliminado correctamente", paramInt)))

	}
}

// UpdateNameAndPrice godoc
// @Summary     Update name and price
// @Tags     Products
// @Description Update name and price
// @Accept      json
// @Produce     json
// @Param       id      path     int          true "Id product"
// @Param       token   header   string       true "TOKEN"
// @Param       product body     Request      true "Product to update"
// @Success     200     {object} web.Response "List Products"
// @Failure     400     {object} web.Response "Bad Request"
// @Failure     401     {object} web.Response "Unauthorized"
// @Failure     404     {object} web.Response "Not found products"
// @Router      /products/{id} [PATCH]
func (s *Product) UpdateNameAndPrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "Token invalido"))
			return
		}

		idParam := ctx.Param("id")
		paramInt, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		var req Request
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		if req.Name == "" {

			ctx.JSON(400, web.NewResponse(400, nil, "El nombre del producto es necesario"))
			return
		}

		if req.Price == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "El precio del producto es necesario"))
			return

		}

		prod, err := s.service.UpdateNameAndPrice(paramInt, req.Name, req.Price)
		if err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, prod, ""))
	}
}
