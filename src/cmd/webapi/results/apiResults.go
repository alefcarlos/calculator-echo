package results

import (
	"net/http"

	"github.com/labstack/echo"
)

//APIResult modelo de resultado API OK
type APIResult struct {
	Result interface{} `json:"result"`
}

//SendJSONSuccess -
func SendJSONSuccess(c echo.Context, i interface{}) error {
	result := &APIResult{
		Result: i,
	}

	return c.JSON(http.StatusOK, result)
}

//SendJSONBadRequest -
func SendJSONBadRequest(c echo.Context, msg string) error {
	result := &APIResult{
		Result: msg,
	}

	return c.JSON(http.StatusBadRequest, result)
}
