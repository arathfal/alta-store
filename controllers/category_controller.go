package controllers

import (
	"AltaStore/configs"
	"AltaStore/models/categories"
	"net/http"

	"github.com/labstack/echo"
)

func GetCategoryController(e echo.Context) error {
	var dataCategory []categories.Category

	err := configs.DB.Find(&dataCategory).Error

	if err != nil {
		return e.JSON(http.StatusInternalServerError, categories.ResponseCategory{
			false, "Failed get category", nil,
		})
	}

	return e.JSON(http.StatusOK, categories.ResponseCategory{
		true, "Success get category", dataCategory,
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
		return e.JSON(http.StatusInternalServerError, categories.ResponseCategory{
			false, "Failed add category", nil,
		})
	}

	return e.JSON(http.StatusOK, categories.ResponseCategorySingle{
		true, "Success add category", categoryDB,
	})
}
