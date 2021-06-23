package routes

import (
	"AltaStore/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/customers", controllers.GetCustomerController)
	e.POST("/customers/register", controllers.RegisterController)
	e.DELETE("carts/:cartId/products/:productsId", controllers.DeleteCartControllers)
	return e
}
