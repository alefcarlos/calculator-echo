package handlers

import (
	"github.com/alefcarlos/calculator-echo/src/cmd/webapi/results"
	"github.com/alefcarlos/calculator-echo/src/internal/repos"
	"github.com/labstack/echo"
)

//GetRecipes handler de health check
func GetRecipes(c echo.Context) error {
	recipes := repos.GetRecipes()
	return results.SendJSONSuccess(c, recipes)
}
