package viewModels

//Calculate modelo com as informações necessárias para cálculo
type Calculate struct {
	SourceCode  string                `json:"sourceCode"`
	Body        interface{}           `json:"body"`
	ExtraValues []CalculateExtraValue `json:"extraValues"`
}

//CalculateExtraValue modelo que contém as informações de valores extras no cálculo
type CalculateExtraValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

//Constantes dos possíveis valores para o campo SourceCode do modelo Calculate
const (
	SourceCodeV1POS = "v1/pos/purchases" //Chamada a partir da POS v1V
)

//SourceCodesAllowed Lista de códigos de modelos que podemos fazer cálculos
var SourceCodesAllowed = []string{SourceCodeV1POS}

//SourceCodeValidate válida se o valor de SourceCode contém na lista de permitidos
func (item *Calculate) SourceCodeValidate() bool {
	return Exists(SourceCodesAllowed, item.SourceCode)
}

//Exists verifica se existe um elemento num array
func Exists(s []string, value string) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}

	return false
}
