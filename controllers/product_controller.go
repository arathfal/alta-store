package controllers

import (
	"AltaStore/configs"
	"AltaStore/models/product"
	"net/http"

	"github.com/labstack/echo"
)

func GetProductsController(e echo.Context) error {
	var dataProduct []product.Product

	err := configs.DB.Preload("Category").Find(&dataProduct).Error

	if err != nil {
		return e.JSON(http.StatusInternalServerError, product.ProductResponse{
			false, "Failed get data products", nil,
		})
	}

	return e.JSON(http.StatusOK, product.ProductResponse{
		true, "Success get data products", dataProduct,
	})
}

func GetProductsByCategoryController(e echo.Context) error {
	var dataProduct []product.Product
	categoryId := e.QueryParam("CategoryID")

	err := configs.DB.Preload("Category").Find(&dataProduct).Where("category_id = ?", categoryId).Error

	if err != nil {
		return e.JSON(http.StatusInternalServerError, product.ProductResponse{
			false, "Failed get data products", nil,
		})
	}

	return e.JSON(http.StatusOK, product.ProductResponse{
		true, "Success get data products", dataProduct,
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
		return e.JSON(http.StatusInternalServerError, product.ProductResponse{
			false, "Failed create products", nil,
		})
	}

	return e.JSON(http.StatusOK, product.ProductResponseSingle{
		true, "Success cteate products", productDB,
	})
}
