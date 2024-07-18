package routes

import (
	"e-commer/infrastructure/entrypoint/server/routes/orders"
	"e-commer/infrastructure/entrypoint/server/routes/products"
	productUseCase "e-commer/pkg/useCases/products"

	"github.com/labstack/echo/v4"
)

func Init(r *echo.Echo) {

	handler := productUseCase.Handler{}

	handler.NewRecordsHandler()

	products.InitPost(r.Group("/products"), handler)

	orders.Init(r.Group("/orders"))
}
