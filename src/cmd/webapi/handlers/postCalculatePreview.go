package handlers

import (
	"encoding/json"

	"github.com/alefcarlos/calculator-echo/src/cmd/webapi/results"
	"github.com/alefcarlos/calculator-echo/src/internal/errors"
	"github.com/alefcarlos/calculator-echo/src/internal/services"
	"github.com/alefcarlos/calculator-echo/src/internal/viewModels"
	"github.com/labstack/echo"
)

//PostCalculationPreview exibe os materiais achados em uma determinada requisição
func PostCalculationPreview(c echo.Context) (err error) {
	var calc = new(viewModels.Calculate)

	if err = c.Bind(calc); err != nil {
		results.SendJSONBadRequest(c, err.Error())
		return
	}

	//Válidar se o source code é permitido para fazer cálculo
	if !calc.SourceCodeValidate() {
		results.SendJSONBadRequest(c, errors.ErrSourceCodeIsInvalid.Error())
		return
	}

	b := []byte{}

	if b, err = json.Marshal(calc.Body); err != nil {
		results.SendJSONBadRequest(c, err.Error())
		return
	}

	//Obter o corpo de acordo com sourcecode
	calcBody, err := services.GetCalcBody(calc.SourceCode, b)
	if err != nil {
		results.SendJSONBadRequest(c, err.Error())
		return
	}

	materials := calcBody.GetMaterials(calc.ExtraValues)

	// Realizar o cálculo
	result, err := services.GetExpressionResultWithPreview(calcBody, materials)
	if err != nil {
		results.SendJSONBadRequest(c, err.Error())
		return
	}

	return results.SendJSONSuccess(c, result)
}
