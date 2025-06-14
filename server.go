package main

import (
	"Zadanie4/controllers"
	"Zadanie4/database"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	database.InitDB()

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"}, // frontend Reactowy
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<h1>Hello, World!</h1>")
	})

	// Products
	const productPath = "/products/:id"
	e.POST("/products", controllers.CreateProduct)
	e.GET("/products", controllers.GetProducts)
	e.GET(productPath, controllers.GetProduct)
	e.PUT(productPath, controllers.UpdateProduct)
	e.DELETE(productPath, controllers.DeleteProduct)

	//Cart
	e.GET("/cart/:id", controllers.GetCart)
	e.POST("/cart", controllers.AddCart)

	e.Logger.Fatal(e.Start(":1323"))

}
