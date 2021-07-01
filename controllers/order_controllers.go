package controllers

import (
	"AltaStore/configs"
	"AltaStore/models"
	// "AltaStore/models/cart"
	"AltaStore/models/customer"
	"AltaStore/models/order"
	"AltaStore/models/orderitems"

	// "AltaStore/models/product"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func PostOrderController(c echo.Context) error {
	var orderCreate order.Order
	c.Bind(&orderCreate)

	var customer customer.Customer
	fmt.Println(orderCreate.CustomerID)
	configs.DB.Where("id = ?", orderCreate.CustomerID).Preload("Cart.CartItems.Product").Find(&customer)
	fmt.Println(customer.Cart.CartItems)

	// if cart customer kosong -- FALSE
	if customer.Cart.CartItems == nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			false, "cart kosong",
		})
		// return echo.NewHTTPError(http.StatusBadRequest, "data cart kosong")
	}

	var orderDB order.Order
	orderDB.StatusPayment = false
	// orderDB.TotalAmount = orderCreate.TotalAmount
	orderDB.OrderDate = time.Now()
	// orderDB.Address = orderCreate.Address
	orderDB.StatusDeliver = "Menunggu Verifikasi Pembayaran"
	// orderDB.PhoneNumber = orderCreate.PhoneNumber
	// orderDB.Payment_Method = orderCreate.Payment_Method

	// get payment method kode pembayaran

	for _, v := range customer.Cart.CartItems {
		orderDB.OrderItems = append(orderDB.OrderItems, orderitems.OrderItems{
			Name: v.Product.Name,
			Quantity: v.Quantity,
			Price: v.Product.Price,
			Description: v.Product.Description,
		})	
	}
	orderDB.CustomerID = orderCreate.CustomerID

	configs.DB.Create(&orderDB)

	// jika error create maka tidak di delete
	
	configs.DB.Where("cart_id = ?", customer.Cart.ID).Delete(&customer.Cart.CartItems)

	return c.JSON(http.StatusOK, models.Response{
		true, "Success Checkout",
	})
}

func PayController(c echo.Context) error {
	var orderUpdate order.Order
	c.Bind(&orderUpdate)


	var order order.Order
	// var customer customer.Customer
	configs.DB.Where("id = ?", orderUpdate.CustomerID).Find(&order)
	// fmt.Println(customer.Cart.CartItems)

	// if cart customer kosong -- FALSE 
	order.StatusPayment = true
	order.StatusDeliver = "Approved"

	configs.DB.Save(&order)

	return c.JSON(http.StatusOK, models.Response{
		true, "Transaction success",
	})
}