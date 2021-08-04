package controllers

import (
	"AltaStore/configs"
	"AltaStore/models"
	"AltaStore/models/categories"
	"net/http"

	"github.com/labstack/echo"
)

func GetCategoryController(e echo.Context) error {
	var dataCategory []categories.Category

	err := configs.DB.Find(&dataCategory).Error

	if err != nil {
		return e.JSON(http.StatusInternalServerError, models.Response{
			Status: false, Message: "Failed Get Category",
		})
	}
	success := models.Response{
		Status: true, Message: "Success Get Category",
	}
	return e.JSON(http.StatusOK, categories.ResponseCategory{
		Response: success, Data: dataCategory,
	})

}

func CreateCategoryController(e echo.Context) error {
	var categoryCreate categories.Category
	e.Bind(&categoryCreate)

	var categoryDB categories.Category
	categoryDB.Name = categoryCreate.Name

	// err := configs.DB.Create(&categoryCreate).Error
	err := configs.DB.Create(&categoryDB).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, models.Response{
			Status: false, Message: "Failed Add Category",
		})
	}
	success := models.Response{
		Status: true, Message: "Success Add Category",
	}
	return e.JSON(http.StatusOK, categories.ResponseCategorySingle{
		Response: success, Data: categoryDB,
	})
}
