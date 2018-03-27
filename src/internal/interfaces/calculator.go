package interfaces

import (
	"github.com/alefcarlos/calculator-echo/src/internal/viewModels"
)

//Calculator representa um modelo calculável
type Calculator interface {
	GetMaterials(extraValues []viewModels.CalculateExtraValue) map[string]interface{}
	GetExtraItemsMaterials(bodyMaterials map[string]interface{}) []*viewModels.ExtraItem
}
