package routes

import (
	"AltaStore/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/customers", controllers.GetCustomerController)
	e.POST("/customers/register", controllers.RegisterController)
	e.GET("/category", controllers.GetCategoryController)
	e.POST("/category", controllers.CreateCategoryController)
	e.DELETE("carts/:cartId/products/:productId", controllers.DeleteCartControllers)
	
	return e
}
