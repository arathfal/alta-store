package controllers

import (
	"AltaStore/configs"
	"AltaStore/models/categories"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var mockDBCategorySucess = categories.Category{
	Name:     "Mainan",
}

func AddCategoryData() bool {
	category := categories.Category{Name: "Test Category"}
	err := configs.DB.Create(&category)
	if err != nil {
		return false
	}
	return false
}

func TestGetCategoryControllers(t *testing.T) {
	e := SetupEchoDB()
	configs.DB.Migrator().DropTable(&categories.Category{})
	configs.DB.AutoMigrate(&categories.Category{})
	AddCategoryData()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/category")

	// Assertions
	if assert.NoError(t, GetCategoryController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		var responseCategory categories.ResponseCategory
		json.Unmarshal([]byte(body), &responseCategory)

		assert.Equal(t, responseCategory.Status, true)
		assert.Equal(t, len(responseCategory.Data), 1)
		assert.Equal(t, responseCategory.Data[0].Name, "Test Category")
	}
	
}

func TestPostCategoriesController(t *testing.T){
	e := SetupEchoDB()
	configs.DB.Migrator().DropTable(&categories.Category{})
	configs.DB.Migrator().AutoMigrate(&categories.Category{})
	// customer := customer.Customer{
		
	// }
	// var json = []byte(`{
  //   "name": "baju"
	// }`)

	body, _ := json.Marshal(&mockDBCategorySucess)

	req := httptest.NewRequest(http.MethodPost, "/categories", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/categories")

	if assert.NoError(t, CreateCategoryController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)	
		body := rec.Body.String()
		var responseCategory categories.ResponseCategory
		fmt.Println(body)
		json.Unmarshal([]byte(body), &responseCategory)

		assert.Equal(t, true, responseCategory.Status)
		assert.Equal(t, "Success Add Category", responseCategory.Message)
	}
	
}