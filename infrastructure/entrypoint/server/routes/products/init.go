package products

import (
	"e-commer/pkg/useCases/products"
	productUseCase "e-commer/pkg/useCases/products"

	"github.com/labstack/echo/v4"
)

type PostRoute struct {
	Handler products.Handler
}

func InitPost(g *echo.Group, h productUseCase.Handler) {
	var p PostRoute

	p.Handler = h

	g.POST("/new", p.CreateProduct)
}
