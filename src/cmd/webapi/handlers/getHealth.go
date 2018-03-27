package handlers

import (
	"github.com/alefcarlos/calculator-echo/src/cmd/webapi/results"
	"github.com/labstack/echo"
)

//GetHealth handler de health check
func GetHealth(c echo.Context) error {
	return results.SendJSONSuccess(c, "tudo ok, patrão!")
}
