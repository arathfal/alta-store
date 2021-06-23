package routes

import (
	"AltaStore/controllers"

	"github.com/labstack/echo"
)


func New() *echo.Echo {
	e := echo.New()
	e.DELETE("carts/:cartId/products/:productsId", controllers.DeleteCartControllers)
	return e
}