package controllers

import (
	"AltaStore/configs"
	"AltaStore/models/categories"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

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
