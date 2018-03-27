package entities

//Recipe modelo contendo informações de como válidar uma formula
type Recipe struct {
	ID          int    `json:"_id,omitempty"`
	Name        string `json:"name"`
	DisplayText string `json:"displayText" `
	Formula     string `json:"formula" `

	//FindFormula regra de devemos utilizar para encontrar melhor formula para uma determinada requisição
	FindFormula string `json:"findFormula"`

	//Priority indica a prioridade da formula
	//Caso seja encontrada mais de uma formula para uma requisição
	//a com maior prioridade será a escolhida
	Priority int `json:"priority"`

	//IsExtra indica se essa fórmula é exclusiva para Dotz Extra
	IsExtra bool `json:"isExtra" `
}

// RecipeByPriority implements sort.Interface for []Recipe based on
// the Priority field.
//https://golang.org/pkg/sort/
type RecipeByPriority []*Recipe

func (a RecipeByPriority) Len() int               { return len(a) }
func (a RecipeByPriority) Swap(i, j int)          { a[i], a[j] = a[j], a[i] }
func (a RecipeByPriority) Less(i int, j int) bool { return a[i].Priority > a[j].Priority }
