package main

import (
	"github.com/alefcarlos/calculator-echo/src/cmd/webapi/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	e.GET("/", handlers.GetHealth)

	e.POST("/recipes", handlers.PostRecipes)
	e.GET("/recipes", handlers.GetRecipes)

	e.POST("/calculate/preview", handlers.PostCalculationPreview)

	e.Logger.Fatal(e.Start(":1323"))
}
