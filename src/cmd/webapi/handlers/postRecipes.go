package handlers

import (
	"github.com/alefcarlos/calculator-echo/src/cmd/webapi/results"
	"github.com/alefcarlos/calculator-echo/src/internal/entities"
	"github.com/alefcarlos/calculator-echo/src/internal/repos"
	"github.com/labstack/echo"
)

//PostRecipes handler do post recipe
func PostRecipes(c echo.Context) (err error) {
	r := new(entities.Recipe)
	if err = c.Bind(r); err != nil {
		return
	}

	repos.AddRecipe(r)

	return results.SendJSONSuccess(c, "tudo ok, patrão!")
}
