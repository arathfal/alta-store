package controllers

import (
	// "AltaStore/configs"
	// "AltaStore/models/product"
	"AltaStore/configs"
	"AltaStore/models"
	"AltaStore/models/cart"
	"AltaStore/models/cartitems"
	"net/http"

	// "net/http"
	"strconv"

	"github.com/labstack/echo"
)

func AddToCartController(c echo.Context) error {
	cartId, _ := strconv.Atoi(c.FormValue("cartId"))
	productId, _ := strconv.Atoi(c.FormValue("productId"))
	quantity, _ := strconv.Atoi(c.FormValue("quantity"))

	var dataCart cart.Cart
	var dataCartItems cartitems.CartItems

	dataCartItems.CartID = cartId
	dataCartItems.ProductID = productId
	dataCartItems.Quantity = quantity

	err := configs.DB.Create(&dataCartItems).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Status: false, Message: "Failed Add to Cart",
		})
	}

	configs.DB.Preload("Products").First(&dataCart, cartId)

	success := models.Response{
		Status: true, Message: "Success Add to Cart",
	}

	return c.JSON(http.StatusOK, cart.ResponseCartSingle{
		Response: success, Data: dataCart,
	})
}

func GetCartControllers(c echo.Context) error {
	var dataCart []cart.Cart

	err := configs.DB.Preload("Products").Find(&dataCart).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Status: false, Message: "Failed Get Cart List",
		})
	}
	success := models.Response{
		Status: true, Message: "Success Get Cart List",
	}

	return c.JSON(http.StatusOK, cart.ResponseCart{
		Response: success, Data: dataCart,
	})
}

func DeleteCartControllers(c echo.Context) error {
	cartId, _ := strconv.Atoi(c.Param("cartId"))
	productId, _ := strconv.Atoi(c.Param("productId"))

	var dataCart cart.Cart
	var dataCartItems cartitems.CartItems

	err := configs.DB.Where("cart_id = ? AND product_id = ?", cartId, productId).Delete(&dataCartItems).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Status: false, Message: "Failed to Delete Data",
		})
	}

	configs.DB.Preload("Products").First(&dataCart, cartId)

	// _ = id
	return c.JSON(http.StatusOK, models.Response{
		Status: true, Message: "Data Deleted",
	})
}
