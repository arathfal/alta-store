package controllers

import (
	"AltaStore/configs"
	"AltaStore/models/product"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func AddProductData() bool {
	product := product.Product{Name: "Test Product", Stock: 1, Price: 1000, Description: "Decs", CategoryID: 1}
	err := configs.DB.Create(&product)
	if err != nil {
		return false
	}
	return false
}

func TestGetProductControllers(t *testing.T) {
	e := SetupEchoDB()
	configs.DB.Migrator().DropTable(&product.Product{})
	configs.DB.AutoMigrate(&product.Product{})
	AddProductData()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/products")

	// Assertions
	if assert.NoError(t, GetProductsController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		var responseProduct product.ProductResponse
		json.Unmarshal([]byte(body), &responseProduct)

		assert.Equal(t, responseProduct.Status, true)
		assert.Equal(t, len(responseProduct.Data), 1)
		assert.Equal(t, responseProduct.Data[0].Name, "Test Product")
	}
}
