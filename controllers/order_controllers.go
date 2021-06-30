package controllers

import (
	"AltaStore/configs"
	"AltaStore/models"
	"AltaStore/models/customer"
	"AltaStore/models/order"
	"AltaStore/models/orderitems"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func PostOrderController(c echo.Context) error {
	var orderCreate order.Order
	c.Bind(&orderCreate)

	var customer customer.Customer
	// var cart cart.Cart
	// var products []product.Product
	fmt.Println(orderCreate.CustomerID)
	configs.DB.Where("id = ?", orderCreate.CustomerID).Preload("Cart.CartItems.Product").Find(&customer)
	fmt.Println(customer.Cart.CartItems)
	// var cart cart.Cart
	// cart := customer.Cart
	// configs.DB.Model(&cart).Association("Products").Find(&products)

	// fmt.Println(products)
	var orderDB order.Order
	orderDB.StatusPayment = false
	// orderDB.TotalAmount = orderCreate.TotalAmount
	orderDB.OrderDate = time.Now()
	// orderDB.Address = orderCreate.Address
	orderDB.StatusDeliver = "Menunggu Verifikasi Pembayaran"
	// orderDB.PhoneNumber = orderCreate.PhoneNumber
	// orderDB.Payment_Method = orderCreate.Payment_Method
	for _, v := range customer.Cart.CartItems {
		orderDB.OrderItems = append(orderDB.OrderItems, orderitems.OrderItems{
			Name:        v.Product.Name,
			Quantity:    v.Quantity,
			Price:       v.Product.Price,
			Description: v.Product.Description,
		})
		// fmt.Println(customer.Cart.CartItems[i])
	}
	orderDB.CustomerID = orderCreate.CustomerID

	configs.DB.Create(&orderDB)

	return c.JSON(http.StatusOK, models.Response{
		false, "Data Created",
	})
}
