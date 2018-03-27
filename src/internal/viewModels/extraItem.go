package viewModels

//ExtraItem modelo com as informações de itens extra
type ExtraItem struct {
	Sku       string
	Materials map[string]interface{}
}

//AddMaterial adiciona um novo material na lista
func (item *ExtraItem) AddMaterial(key string, value interface{}) {
	if item.Materials == nil {
		item.Materials = make(map[string]interface{}, 0)
	}

	item.Materials[key] = value
}
