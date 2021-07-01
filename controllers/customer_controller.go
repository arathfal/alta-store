package controllers

import (
	"AltaStore/configs"
	helper "AltaStore/helpers"
	"AltaStore/middlewares"
	"AltaStore/models"
	"AltaStore/models/customer"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func GetCustomerController(e echo.Context) error {
	var dataCustomers []customer.Customer

	err := configs.DB.Find(&dataCustomers).Error

	if err != nil {
		return e.JSON(http.StatusInternalServerError, models.Response{
			Status: false, Message: "Failed Get Data Customers",
		})
	}

	success := models.Response{
		Status: true, Message: "Success Get Data Customers",
	}

	return e.JSON(http.StatusOK, customer.ResponseCustomer{
		Response: success, Data: dataCustomers,
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
		return e.JSON(http.StatusInternalServerError, models.Response{
			Status: false, Message: "Failed Register",
		})
	}
	success := models.Response{
		Status: true, Message: "Success Registered",
	}

	return e.JSON(http.StatusOK, customer.ResponseCustomerSingle{
		Response: success, Data: customerDB,
	})
}

func CheckLogin(email, password string) (bool, int, error) {
	var customerDB customer.Customer
	var customer customer.Customer

	err := configs.DB.Debug().Where("email = ?", email).Find(&customerDB).Scan(&customer).Error
	id := customer.ID
	pwd := customer.Password

	if err != nil {
		fmt.Println("Email not found")
		return false, 0, err
	}

	match, err := helper.CheckHashAndPass(password, pwd)

	if !match {
		fmt.Println("Password doesn't match")
		return false, 0, err

	}
	return true, int(id), nil

}

func LoginController(e echo.Context) error {
	email := e.FormValue("Email")
	password := e.FormValue("Password")
	result, id, err := CheckLogin(email, password)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, models.Response{
			Status: false, Message: "Password Invalid",
		})
	}

	if !result {
		return echo.ErrUnauthorized
	}

	token, _ := middlewares.GenerateToken(id, email)

	success := models.Response{
		Status: true, Message: "Success Login, Welcome to Alta Store",
	}

	return e.JSON(http.StatusOK, customer.LoginResponse{
		Response: success, Token: token,
	})
}
