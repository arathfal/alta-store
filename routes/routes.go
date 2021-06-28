package routes

import (
	"AltaStore/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/products", controllers.GetProductsController)
	e.POST("/products", controllers.CreateProductController)
	return e
}
