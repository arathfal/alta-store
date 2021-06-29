package routes

import (
	"AltaStore/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/customers", controllers.GetCustomerController)
	e.POST("/customers/register", controllers.RegisterController)
	e.GET("/products", controllers.GetProductsController)
	e.GET("/products?category=id", controllers.GetProductsByCategoryController)
	e.POST("/products", controllers.CreateProductController)
	e.GET("/category", controllers.GetCategoryController)
	e.POST("/category", controllers.CreateCategoryController)
	return e
}
