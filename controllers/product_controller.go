package controllers

import (
	"Zadanie4/database"
	"Zadanie4/models"
	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateProduct(c echo.Context) error {
	var product models.Product
	bindError := c.Bind(&product)
	if bindError == nil {		
		var newProduct = models.Product{
			Name: product.Name,
			Price: product.Price,
		}
		res := database.DB.Create(&newProduct)
		if res.Error != nil {
			return c.JSON(http.StatusBadRequest, res.Error)
		}
		return c.JSON(http.StatusOK, newProduct)
	} else{
		return c.JSON(http.StatusBadRequest, bindError)
	}
	
}

func GetProduct(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID format")
	}

	var product models.Product
	result2 := database.DB.First(&product, uint(id))
	if result2.Error != nil {
		return c.String(http.StatusNotFound, "Product not found")
	}
	return c.JSON(http.StatusOK, product)
}

func GetProducts(c echo.Context) error {
	var products []models.Product
	result := database.DB.Find(&products)
	if result.Error != nil {
		return c.String(http.StatusNotFound, "No products found")
	}
  	return c.JSON(http.StatusOK, products)
}

func UpdateProduct(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID format")
	}
	
	var product models.Product
	var updatedProduct models.Product
	bindError := c.Bind(&updatedProduct)
	if bindError == nil {
		res := database.DB.Find(&product, uint(id))
		if res.Error != nil {
			return c.String(http.StatusNotFound, "Product not found")
		} else {
			database.DB.Model(&product).Updates(updatedProduct)
			return c.JSON(http.StatusOK, updatedProduct)
		}
	} else {
		return c.JSON(http.StatusBadRequest, bindError)
	}
}

func DeleteProduct (c echo.Context) error {
	id := c.Param("id")
	var product models.Product
	result := database.DB.Delete(&product, id)
	if result.Error != nil {
		return c.String(http.StatusNotFound, "Couldn't find product with id: "+id)
	}
	return c.String(http.StatusOK, "Deleted product with id: "+id)
}
