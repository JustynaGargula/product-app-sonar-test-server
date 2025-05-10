package controllers

import (
	"Zadanie4/database"
	"Zadanie4/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetCart(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID format")
	}

	var cart models.Cart
	result2 := database.DB2.Preload("Items").First(&cart, uint(id))
	if result2.Error != nil {
		return c.String(http.StatusNotFound, "Cart not found")
	}
	return c.JSON(http.StatusOK, cart)
}

func AddCart(c echo.Context) error {
	var req struct {
		Items []models.CartItem `json:"items"`
	}
	bindError := c.Bind(&req)

	if bindError == nil {
		cart := models.Cart{}
		if err := database.DB2.Create(&cart).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create cart"})
		}

		for _, item := range req.Items {
			item.CartID = cart.ID
			if err := database.DB2.Create(&item).Error; err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create cart item"})
			}
		}
		return c.JSON(http.StatusOK, cart)
	} else {
		return c.JSON(http.StatusBadRequest, bindError)
	}
}
