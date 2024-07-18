package products

import (
	"e-commer/pkg/domain/products"

	"github.com/labstack/echo/v4"
)

func (p *PostRoute) CreateProduct(c echo.Context) error {
	domain := products.CreateProductRequest{}
	if err := c.Bind(&domain); err != nil {
		return err
	}
	response := p.Handler.Post.CreateProduct(domain)

	return c.JSON(200, response)
}
