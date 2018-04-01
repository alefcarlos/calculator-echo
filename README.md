# Calculadora de modelos em GO

## Apredizado

Projeto criado utilizando o framework Echo penasndo no aprendizado.

## Dependêncas

Utilizo o `dep` como gerenciador de pacotes.

>[Baixe](https://github.com/golang/dep/releases) o binário para seu SO no seu $GOROOT/bin

## Rodar local

### Rodar usando go run webapi
```
dep ensure
go run src/cmd/webapi/main.go
```

##  Rodar usando docker

### Clear compose cache(caso necessário)
```cmd
docker-compose rm
```

### Builder docker-compose
```cmd
docker-compose up --build
```

>As vezes dá uns paus, então rodar por partes:

```cmd
docker-compose build --no-cache
docker-compose up
```


# Como usar

## Cadastrar fórmulas

> Fórmula é a maneira de cálculo de uma determinada requisição

Encontramos a melhor fórmula através da resolução do campo `findFormula`, caso duas ou mais fórmulas sejam encontradas, devemos pegar a com maior prioridade.

`POST /recipes` passando o seguinte JSON:

```json
{
 "name": "Formula Ganhe Online",
 "displayText": "DZ 5 por 1 R$",
 "findFormula": "(sponsor == '4334343' && date > '2018-02-01' && date < '2018-02-28 23:59:59')",
 "formula":"purchaseValue * 5",
 "priority": 1,
 "isExtra": false,
 "usedMaterials": {"sponsorId": "4334343"}
}
``` 

## Testar fórmulas

> Obtemos as variáveis da requisição a partir do mapeamento da mesma

`POST /calculate/preview` passando o seguinte JSON:

```json
{
  "sourceCode": "v1/pos/purchases",
  "body": {
    "locator": "10",
    "storeCode": "MATRIZ",
    "deviceCode": "41",
    "purchaseDate": "2018-02-09T19:39:43.564Z",
    "shopperIdentification": "43574989881",
    "employeeIdentification": "43574989881",
    "calculatePoints": false,
    "offerCode": "GDkqlEVZREMw0pj",
    "points": 10,
    "items": [
      {
        "sku": "132312312",
        "ean": "0768421222301",
        "quantity": 2,
        "unitPrice": 8.0,
        "itemPrice": 16.0,
        "supplierIdentification": "43824418000126",
        "name": "ARROZ INTEGRAL",
        "discount": 0.0,
        "offerCode": "GDkqlEVZREMw0pj",
        "points": 1
      }
    ],
    "purchaseValue": 560,
    "payments": [
      {
        "paymentType": 1,
        "bin": "",
        "amount": 16.00
      }
    ]
  },
  "extraValues": [
    {
      "key": "sponsor",
      "value": "4334343"
    }
  ]
}
```

# Modelos de cálculos

> São os modelos mapeados com a maneira de como extrair variáveis

Todos os modelos de cálculos devem ser criados na pasta `src/internal/calculationModels` e devem implementar a interface `interfaces.Calculator`

```golang
GetMaterials(extraValues []viewModels.CalculateExtraValue) map[string]interface{}
GetExtraItemsMaterials(bodyMaterials map[string]interface{}) []*viewModels.ExtraItem
```

> `GetMaterials` retorna todas as variáveis contidas no modelo

> `GetExtraItemsMaterials` retorna todas as variáeis contidas nos itens extras de uma requisição

> **Devemos também criar uma cosntante contendo a identificação do modelo(SourceCod) em `src/internal/viewModels/calculate.go`**

## Modelos existentes

* v1/pos/purchase

## v1/pos/purchase

JSON representando a requisição

```json
{
    "locator": "10",
    "storeCode": "MATRIZ",
    "deviceCode": "41",
    "purchaseDate": "2018-02-09T19:39:43.564Z",
    "shopperIdentification": "43574989881",
    "employeeIdentification": "43574989881",
    "calculatePoints": false,
    "offerCode": "GDkqlEVZREMw0pj",
    "points": 10,
    "items": [
      {
        "sku": "132312312",
        "ean": "0768421222301",
        "quantity": 2,
        "unitPrice": 8.0,
        "itemPrice": 16.0,
        "supplierIdentification": "43824418000126",
        "name": "ARROZ INTEGRAL",
        "discount": 0.0,
        "offerCode": "GDkqlEVZREMw0pj",
        "points": 1
      }
    ],
    "purchaseValue": 560,
    "payments": [
      {
        "paymentType": 1,
        "bin": "",
        "amount": 16.00
      }
    ]
  }
```

## Variáveis

As variáveis disponíveis para uso são:

### Base

* date - Data da compra
* purchaseValue - Valor total da compra
* storeCode - Identificação da Loja
* deviceCode - Identificação PDV

### Itens extra

> Retorna uma lista com os campos abaixo, cada item terá essas informações

* sku - Identificação do item
* unitPrice - Valor unitário do item
* quantity - Quantidade comprada
* totalItemPrice - quantity * unitPrice

# Fórmulas

Um JSON representando uma entidade de fórmula

```json
{
    "name": "Formula Por loja",
    "displayText": "DZ 15 por 1 R$(storeCode)",
    "findFormula": "(sponsor == '4334343' && storeCode in ('MATRIZ') && date > '2018-02-01' && date < '2018-02-29 23:59:59')",
    "formula":"purchaseValue * 15",
    "priority": 2,
    "isExtra": false,
    "usedMaterials": {"sponsorId": "4334343"}
}
```

## Explicação dos campos

`name` - Identificação da fórmula

`displayText` - Descrição de como apareceria no extrato do cliente

`findFormula` - Fórmula de match com a requisição utilizando variáveis, o resultado esperado é um boolean

`formula` - Fórmula de cálculo, pode-se usar variáveis ou somente valor fixo, exemplo: purchaseValue * 5 ou 1500, o resultado esperado é um float64

`priority` - Indica prioridade das fórmula, quando encontrada duas através do findFormula, a com maior será a escolhida

`isExtra` - Indica se essa fórmula é específica para Cálculos de Dotz Extras

`usedMaterials` - Indica a lista de variáveis utilizadas na fórmula