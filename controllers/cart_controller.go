package controllers

import (
	// "AltaStore/configs"
	// "AltaStore/models/product"
	"AltaStore/configs"
	"AltaStore/models/cart"
	"AltaStore/models/cartitems"
	"AltaStore/models/product"
	"fmt"
	"net/http"

	// "net/http"
	"strconv"

	"github.com/labstack/echo"
)


func DeleteCartControllers(c echo.Context) error {
	cartId, err := strconv.Atoi(c.Param("cartId"))
	fmt.Println(cartId)
	productId, _ := strconv.Atoi(c.Param("productId"))
	fmt.Println(productId)

	var dataProducts []product.Product
	var dataCart cart.Cart
	var dataCartItems cartitems.CartItems

	// err = configs.DB.Model(&cart.Cart{}).Association("Products").Find(&dataProducts)
	// err = configs.DB.Debug().Model(&dataCart).Where("cart_id = ?", cartId).Association("Products").Find(&dataProducts)
	// err = configs.DB.Debug().Preload("Products").Find(&dataCart).
	// err = configs.DB.Model(&dataCart).Association("Products").Error
	// if (err != nil) {
	// 	return c.JSON(http.StatusBadRequest, err.Error())
	// }
	// configs.DB.Model(&dataCart).Association("Products").Find(&dataProducts)
	// count := configs.DB.Debug().Model(&dataCart).Association("Products").Count()
	// fmt.Println(count)
	// err = configs.DB.Debug().Model(dataProducts).Where("id = ?", cartId).Association("Carts").Find(&dataCart)
	// configs.DB.Debug().Model(&dataCart).Association("Products").Find(&dataProducts)
	// configs.DB.Debug().Find(&dataProducts)
	// err = configs.DB.Find(&dataCart).Error
	// db.Model(&user).Association("Languages").Delete(languageZH, languageEN)
	fmt.Println(err)
	// db.Model(&user).Association("Languages").Find(&languages)
	// err = configs.DB.Find(&dataProducts).Error
	fmt.Println(dataProducts)
	fmt.Println(dataCart)
	// fmt.Println(dataCart)
	
	configs.DB.Where("cart_id = ? AND product_id = ?", cartId, productId).Delete(&dataCartItems)
	fmt.Println(dataCart)
	
	configs.DB.Preload("Products").First(&dataCart, cartId)
	fmt.Println(dataCart)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, product.Product{
			
		})
	}
	
	// _ = id
	return c.JSON(http.StatusOK, dataCart)
}