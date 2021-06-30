package controllers

import (
	"AltaStore/configs"
	helper "AltaStore/helpers"
	"AltaStore/middlewares"
	"AltaStore/models/customer"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func GetCustomerController(e echo.Context) error {
	var dataCustomers []customer.Customer

	err := configs.DB.Find(&dataCustomers).Error

	if err != nil {
		return e.JSON(http.StatusInternalServerError, customer.ResponseCustomer{
			false, "Failed get data Customers", nil,
		})
	}

	return e.JSON(http.StatusOK, customer.ResponseCustomer{
		true, "Success get Data Customer", dataCustomers,
	})
}

func RegisterController(e echo.Context) error {
	var customerRegister customer.CustomerRegister
	e.Bind(&customerRegister)

	var customerDB customer.Customer
	customerDB.Name = customerRegister.Name
	customerDB.Email = customerRegister.Email
	customerDB.Password = helper.HashAndSalt(customerRegister.Password)

	err := configs.DB.Create(&customerDB).Error

	if err != nil {
		return e.JSON(http.StatusInternalServerError, customer.ResponseCustomer{
			false, "Failed register", nil,
		})
	}

	return e.JSON(http.StatusOK, customer.ResponseCustomerSingle{
		true, "Success registered", customerDB,
	})
}

func CheckLogin(email, password string) (bool, int, error) {
	var customer customer.Customer
	var pwd string

	err := configs.DB.Select("password").Where("email = ?", email).Take(&customer).Scan(&pwd).Error

	if err != nil {
		fmt.Println("Email not found")
		return false, 0, err
	}

	match, err := helper.CheckHashAndPass(password, pwd)

	if !match {
		fmt.Println("Password doesn't match")
		return false, 0, err

	}
	return true, int(customer.ID), nil

}

func LoginController(e echo.Context) error {
	email := e.FormValue("Email")
	password := e.FormValue("Password")
	result, id, err := CheckLogin(email, password)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, customer.LoginResponse{
			false, "Password Invalid", "",
		})
	}

	if !result {
		return echo.ErrUnauthorized
	}

	token, err := middlewares.GenerateToken(id, email)

	return e.JSON(http.StatusOK, customer.LoginResponse{
		true, "Success Login, welcome to Alta Store", token,
	})
}
