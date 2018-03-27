package calculationmodels

import (
	"log"
	"time"

	"github.com/alefcarlos/calculator-echo/src/internal/viewModels"
)

//POSv1 modelo que representa requisição vinda da POS
//SourceCode: "v1/pos/purchases"
type POSv1 struct {
	Locator                string    `json:"locator"`
	StoreCode              string    `json:"storeCode"`
	DeviceCode             string    `json:"deviceCode"`
	PurchaseDate           time.Time `json:"purchaseDate"`
	ShopperIdentification  string    `json:"shopperIdentification"`
	EmployeeIdentification string    `json:"employeeIdentification"`
	CalculatePoints        bool      `json:"calculatePoints"`
	OfferCode              string    `json:"offerCode"`
	Points                 int       `json:"points"`
	Items                  []struct {
		Sku                    string  `json:"sku"`
		Ean                    string  `json:"ean"`
		Quantity               int     `json:"quantity"`
		UnitPrice              float64 `json:"unitPrice"`
		ItemPrice              float64 `json:"itemPrice"`
		SupplierIdentification string  `json:"supplierIdentification"`
		Name                   string  `json:"name"`
		Discount               float64 `json:"discount"`
		OfferCode              string  `json:"offerCode"`
		Points                 int     `json:"points"`
	} `json:"items"`
	PurchaseValue float64 `json:"purchaseValue"`
	Payments      []struct {
		PaymentType int     `json:"paymentType"`
		Bin         string  `json:"bin"`
		Amount      float64 `json:"amount"`
	} `json:"payments"`
}

//GetMaterials Obtém a lista de materiais/variáveis contidos no modelo
func (item *POSv1) GetMaterials(extraValues []viewModels.CalculateExtraValue) map[string]interface{} {
	result := make(map[string]interface{}, 0)

	//Verificar se existe algum valor em extraValue, então adicionar no retorno
	for _, extra := range extraValues {
		result[extra.Key] = extra.Value
	}

	//Adicionar valores do modelo
	result["date"] = item.PurchaseDate.Unix()
	result["purchaseValue"] = item.PurchaseValue
	result["storeCode"] = item.StoreCode

	return result
}

//GetExtraItemsMaterials obtém a lista de itens extras junto com a lista de possíveis variáveis
//bodyMaterials lista de variáveis extraídas do corpo
func (item *POSv1) GetExtraItemsMaterials(bodyMaterials map[string]interface{}) []*viewModels.ExtraItem {
	result := []*viewModels.ExtraItem{}

	for _, value := range item.Items {
		extraItem := viewModels.ExtraItem{
			Sku: value.Sku,
		}

		//Adiciona as variáveis da requisição
		for key, value := range bodyMaterials {
			extraItem.AddMaterial(key, value)
		}

		extraItem.AddMaterial("sku", value.Sku)
		extraItem.AddMaterial("unitPrice", value.UnitPrice)
		extraItem.AddMaterial("unitPrice", value.UnitPrice)
		extraItem.AddMaterial("quantity", value.Quantity)
		extraItem.AddMaterial("totalItemPrice", value.ItemPrice)

		log.Printf("Extra item: %v", extraItem)

		result = append(result, &extraItem)
	}

	return result
}
