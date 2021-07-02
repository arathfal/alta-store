package controllers

import (
	"AltaStore/configs"
	helper "AltaStore/helpers"
	"AltaStore/models/cart"
	"AltaStore/models/customer"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	mockDBRegisterSucess = customer.Customer{
	Name: "Test Customer",
	Email: "test@gmail.com",
	Password: "test123",
	}
	mockDBLoginSuccess = customer.Customer{
		Name: "Test Customer",
		Email: "test@gmail.com",
		Password: "test123",
	}
)




func AddCustomerData() bool {
	password := "test123"
	pw := helper.HashAndSalt(password)
	fmt.Println(pw)
	user := customer.Customer{Name: "Test Customer", Email: "test@gmail.com", Password: pw}
	err := configs.DB.Create(&user)
	if err != nil {
		return false
	}
	return false
}

func TestGetCustomerControllers(t *testing.T) {
	e := SetupEchoDB()
	configs.DB.Migrator().DropTable(&customer.Customer{})
	configs.DB.AutoMigrate(&customer.Customer{})
	AddCustomerData()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/customers")

	// Assertions
	if assert.NoError(t, GetCustomerController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		var responseCustomer customer.ResponseCustomer
		json.Unmarshal([]byte(body), &responseCustomer)

		assert.Equal(t, responseCustomer.Status, true)
		assert.Equal(t, len(responseCustomer.Data), 1)
		assert.Equal(t, responseCustomer.Data[0].Name, "Test Customer")
	}
}

func TestRegisterController(t *testing.T) {
	e := SetupEchoDB()
	configs.DB.Migrator().DropTable(&customer.Customer{})
	configs.DB.AutoMigrate(&customer.Customer{}, &cart.Cart{})
	// AddCustomerData()
	body, _ := json.Marshal(&mockDBRegisterSucess)
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/register")

	// Assertions
	if assert.NoError(t, RegisterController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		var responseCustomer customer.ResponseCustomerSingle
		json.Unmarshal([]byte(body), &responseCustomer)

		assert.Equal(t, true, responseCustomer.Status)
		assert.Equal(t, "Success Registered", responseCustomer.Message)
		assert.Equal(t, "test@gmail.com", responseCustomer.Data.Email)
	}
}

func TestCheckLoginController(t *testing.T) {
	e := SetupEchoDB()
	configs.DB.Migrator().DropTable(&customer.Customer{})
	configs.DB.AutoMigrate(&customer.Customer{})
	AddCustomerData()
	body, _ := json.Marshal(&mockDBLoginSuccess)
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/login")

	// Assertions
	if assert.NoError(t, LoginController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		var responseCustomer customer.ResponseCustomer
		fmt.Println(body)
		json.Unmarshal([]byte(body), &responseCustomer)

		assert.Equal(t, true, responseCustomer.Status)
		assert.Equal(t, "Success Login, Welcome to Alta Store", responseCustomer.Message)
	}
}