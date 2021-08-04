package controllers

import (
	"AltaStore/configs"
	"AltaStore/models"
	"AltaStore/models/product"
	"net/http"

	"github.com/labstack/echo"
)

func GetProductsController(e echo.Context) error {
	var dataProduct []product.Product

	err := configs.DB.Preload("Category").Find(&dataProduct).Error

	if err != nil {
		return e.JSON(http.StatusInternalServerError, models.Response{
			Status: false, Message: "Failed get data products",
		})
	}

	success := models.Response{
		Status: true, Message: "Success get data products",
	}

	return e.JSON(http.StatusOK, product.ProductResponse{
		Response: success, Data: dataProduct,
	})
}

func GetProductsByCategoryController(e echo.Context) error {
	var dataProduct []product.Product
	categoryId := e.QueryParam("CategoryID")

	err := configs.DB.Preload("Category").Find(&dataProduct).Where("category_id = ?", categoryId).Error

	if err != nil {
		return e.JSON(http.StatusInternalServerError, models.Response{
			Status: false, Message: "Failed get data products",
		})
	}

	success := models.Response{
		Status: true, Message: "Success get data products",
	}

	return e.JSON(http.StatusOK, product.ProductResponse{
		Response: success, Data: dataProduct,
	})

}

func CreateProductController(e echo.Context) error {
	var createProduct product.ProductCreate

	e.Bind(&createProduct)

	var productDB product.Product
	productDB.Name = createProduct.Name
	productDB.Stock = createProduct.Stock
	productDB.Price = createProduct.Price
	productDB.Description = createProduct.Description
	productDB.CategoryID = createProduct.CategoryID

	err := configs.DB.Create(&productDB).Error

	if err != nil {
		return e.JSON(http.StatusInternalServerError, models.Response{
			Status: false, Message: "Failed Add Products",
		})
	}

	success := models.Response{
		Status: true, Message: "Success Add Data Products",
	}

	return e.JSON(http.StatusOK, product.ProductResponseSingle{
		Response: success, Data: productDB,
	})
}
