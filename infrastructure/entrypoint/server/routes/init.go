package routes

import (
	"e-commer/infrastructure/entrypoint/server/routes/orders"
	"e-commer/infrastructure/entrypoint/server/routes/products"

	"github.com/labstack/echo/v4"
)

func Init(r *echo.Echo) {
	products.Init(r.Group("/products"))

	orders.Init(r.Group("/orders"))
}
