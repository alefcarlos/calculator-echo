package viewModels

//CalculateResult modelo para exibição do resultado de um cálculo
type CalculateResult struct {
	Result        float64                     `json:"result"`
	PreviewString string                      `json:"preview,omitempty"`
	DisplayText   string                      `json:"displayText,omitempty"`
	Formula       string                      `json:"formula,omitempty"`
	ExtraItems    []*CalculateResultExtraItem `json:"extraItems,omitempty"`
}

//CalculateResultExtraItem modelo para exibição do resultado de um cálculo de dotz extra
type CalculateResultExtraItem struct {
	SKU           string  `json:"sku"`
	Result        float64 `json:"result"`
	PreviewString string  `json:"preview,omitempty"`
	DisplayText   string  `json:"displayText,omitempty"`
	Formula       string  `json:"formula,omitempty"`
}
