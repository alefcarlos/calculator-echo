# Calculadora de modelos em GO

## Utilizando framework Echo

Projeto crido com intuito de aprender Echo

## Rodar local

### Rodar usando go run webapi
```
dep init
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