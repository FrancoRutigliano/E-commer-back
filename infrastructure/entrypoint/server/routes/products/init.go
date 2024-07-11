package products

import "github.com/labstack/echo/v4"

func Init(g *echo.Group) {
	g.GET("/all", echo.HandlerFunc(func(c echo.Context) error {
		return c.JSON(200, echo.Map{"message": "all products"})
	}))
}
