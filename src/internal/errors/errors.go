package errors

import "errors"

//ErrIDNotFound é um erro que é apresentando quando o ID passado não existe
var ErrIDNotFound = errors.New("O ID informado não foi encontrado")

//ErrMaterialNotSupported é um erro que é apresentado quando o material tem valor fixo
var ErrMaterialNotSupported = errors.New("Esse material não permite adição de valores")

//ErrCouldntGetBodyFromCalc é um erro que é apresentado quando o corpo do cálculo não pôde ser serealizado
var ErrCouldntGetBodyFromCalc = errors.New("Não foi possível realizar o parse do body do cálculo")

//ErrCouldntReadFormula é um erro que é apresentado quando não conseguimos interpretar a formula
var ErrCouldntReadFormula = errors.New("Não foi possível ler a formula")

//ErrCouldntFindRecipe é um erro que é apresentado quando não conseguimos achar a melhor formula para a requisição
var ErrCouldntFindRecipe = errors.New("Não foi possível encontrar formula para a requisição")

//ErrCouldFindRecipes erro é lançado quando não encontra as receitas na base do mongo
var ErrCouldFindRecipes = errors.New("Não foi possível obter as receitas")

//ErrIDIsNotValid erro é lançado quando informamos um ID que não esteja no formato bson.Hex
var ErrIDIsNotValid = errors.New("O ID informado não é válido")

//ErrRequestBodyIsInvalid erro é lançado quando requisitamos um POST com corpo vazio
var ErrRequestBodyIsInvalid = errors.New("Informe um corpo de requisição válido")

//ErrSourceCodeIsInvalid erro é lançado quando requisitamos /calculate ou /calculate/preview com um sourceCode que não esteja presente
//no array viewModels.SourceCodesAllowed
var ErrSourceCodeIsInvalid = errors.New("O Source code não é permitido")
