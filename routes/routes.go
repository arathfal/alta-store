package routes

import (
	"AltaStore/controllers"
	"AltaStore/middlewares"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.POST("/register", controllers.RegisterController)
	e.POST("/login", controllers.LoginController)
	eJwt := e.Group("")
	eJwt.Use(middleware.JWT([]byte(middlewares.GetSecretKey())))
	eJwt.GET("/customers", controllers.GetCustomerController)
	eJwt.GET("/products", controllers.GetProductsController)
	eJwt.GET("/products?category_id=id", controllers.GetProductsByCategoryController)
	eJwt.POST("/products", controllers.CreateProductController)
	eJwt.GET("/category", controllers.GetCategoryController)
	eJwt.POST("/category", controllers.CreateCategoryController)
	eJwt.GET("/carts", controllers.GetCartControllers)
	eJwt.POST("/carts", controllers.AddToCartController)
	eJwt.DELETE("/carts/:cartId/products/:productId", controllers.DeleteCartControllers)
	eJwt.POST("/checkout", controllers.PostOrderController)
	eJwt.POST("/paid", controllers.PayController)

	return e
}
