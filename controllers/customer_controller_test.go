package controllers

import (
	"AltaStore/configs"
	"AltaStore/models/customer"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func AddCustomerData() bool {
	user := customer.Customer{Name: "Test Customer", Email: "test@gmail.com", Password: "test123"}
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
