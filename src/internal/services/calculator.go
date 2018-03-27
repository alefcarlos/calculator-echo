package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"sort"

	"github.com/Knetic/govaluate"
	"github.com/alefcarlos/calculator-echo/src/internal/calculationModels"
	"github.com/alefcarlos/calculator-echo/src/internal/errors"
	"github.com/alefcarlos/calculator-echo/src/internal/repos"

	"github.com/alefcarlos/calculator-echo/src/internal/entities"
	"github.com/alefcarlos/calculator-echo/src/internal/interfaces"
	"github.com/alefcarlos/calculator-echo/src/internal/viewModels"
)

//CalculateExpression calcula o valor de Dotz de acordo com a formula
func CalculateExpression(materials map[string]interface{}, recipe *entities.Recipe) (float64, error) {
	expression, err := govaluate.NewEvaluableExpression(recipe.Formula)
	if err != nil {
		return 0, errors.ErrCouldntReadFormula
	}

	result, err := expression.Evaluate(materials)
	if err != nil {
		return 0, errors.ErrCouldntReadFormula
	}

	return result.(float64), nil
}

//MatchFormula testa uma formula e retorna true
// se o resultado da formula é verdadeira
func MatchFormula(materials map[string]interface{}, formula string) bool {
	expression, err := govaluate.NewEvaluableExpression(formula)
	if err != nil {
		return false
	}

	result, err := expression.Evaluate(materials)
	if err != nil {
		log.Print("Erro ao tentar resolver a fórmula ", err)
		return false
	}

	return result.(bool)
}

//SearchRecipe obtém a melhor formula para a requisição
func SearchRecipe(recipes []*entities.Recipe, materials map[string]interface{}) *entities.Recipe {

	found := repos.FilterRecipe(recipes, func(r *entities.Recipe) bool {
		return MatchFormula(materials, r.FindFormula)
	})

	sort.Sort(entities.RecipeByPriority(found))

	if len(found) == 0 {
		return nil
	}

	return found[0]
}

//GetCalcBody é uma factory para obter o corpo de acordo com sourceCode
func GetCalcBody(sourceCode string, bytes []byte) (interfaces.Calculator, error) {
	switch sourceCode {
	case viewModels.SourceCodeV1POS:
		result := calculationmodels.POSv1{}

		if err := GetPOSv1(bytes, &result); err != nil {
			return nil, err
		}

		return &result, nil
	default:
		return nil, nil
	}
}

//GetExpressionResultValue é uma factory para obter o resultado de uma expressão
func GetExpressionResultValue(materials map[string]interface{}, recipe *entities.Recipe) (float64, error) {
	expressionResult, err := CalculateExpression(materials, recipe)
	if err != nil {
		return 0, err
	}

	return expressionResult, nil
}

//GetExpressionResult obtém o resultado de uma expressão sem preview
func GetExpressionResult(sourceCode string, calcbody interfaces.Calculator, materials map[string]interface{}) (*viewModels.CalculateResult, error) {
	resultModel, recipe, err := calculate(calcbody, materials, false)
	if err != nil {
		return nil, err
	}

	resultModel.DisplayText = recipe.DisplayText

	return resultModel, nil
}

//GetExpressionResultWithPreview é uma factory para obter o resultado de uma expressão
//calcbody modelo de cálculo
//materials variáveis extraídas do body requisição
//retorna também as variáveis e seus valores
func GetExpressionResultWithPreview(calcbody interfaces.Calculator, materials map[string]interface{}) (*viewModels.CalculateResult, error) {
	preview, recipe, err := calculate(calcbody, materials, true)
	if err != nil {
		return nil, err
	}

	preview.DisplayText = recipe.DisplayText
	preview.Formula = recipe.Formula

	//Montar expressão resultante
	var buffer bytes.Buffer
	for key, value := range materials {
		buffer.WriteString(fmt.Sprint("[", key, "] = ", value, ", "))
	}

	preview.PreviewString = buffer.String()

	return preview, nil
}

//GetPOSv1 Obtém o valor do body parseado para o modelo calculationmodels.POSv1
func GetPOSv1(bytes []byte, model *calculationmodels.POSv1) error {
	if err := json.Unmarshal(bytes, &model); err != nil {
		return errors.ErrCouldntGetBodyFromCalc
	}

	return nil
}

func calculateExtraItem(channel chan *viewModels.CalculateResultExtraItem, recipes []*entities.Recipe, extra *viewModels.ExtraItem, isPreview bool) {
	log.Printf("Tentando cálcular item %v", extra)

	recipe := SearchRecipe(recipes, extra.Materials)

	//encontrou fórmula ?
	if recipe == nil {
		log.Print("Não encontrou fórmula para o item extra.")
		channel <- nil
		return
	}

	//A fórmula é exclusiva de dotz extra
	if !recipe.IsExtra {
		log.Print("A fórmula não é compatível com dotz extra.")
		channel <- nil
		return
	}

	log.Printf("Realizando cálculo de dotz extra..")

	//Calcular o item, só será adicionado os itens que conseguimos calcular
	result, err := GetExpressionResultValue(extra.Materials, recipe)
	if err == nil {

		if isPreview {
			//Montar expressão resultante
			var buffer bytes.Buffer
			for key, value := range extra.Materials {
				buffer.WriteString(fmt.Sprint("[", key, "] = ", value, ", "))
			}

			channel <- &viewModels.CalculateResultExtraItem{
				SKU:           extra.Sku,
				Result:        result,
				Formula:       recipe.Formula,
				DisplayText:   recipe.DisplayText,
				PreviewString: buffer.String(),
			}
		} else {
			channel <- &viewModels.CalculateResultExtraItem{
				SKU:         extra.Sku,
				Result:      result,
				DisplayText: recipe.DisplayText,
			}
		}
	}
}

func calculate(calcbody interfaces.Calculator, materials map[string]interface{}, isPreview bool) (*viewModels.CalculateResult, *entities.Recipe, error) {
	resultModel := viewModels.CalculateResult{}

	//Obter todas as recipes
	recipes := repos.GetRecipes()

	//Obter a recipe que se caixa na requisição
	recipe := SearchRecipe(recipes, materials)

	if recipe == nil {
		return nil, nil, errors.ErrCouldntFindRecipe
	}

	result, err := GetExpressionResultValue(materials, recipe)
	if err != nil {
		return nil, nil, err
	}

	resultModel.Result = result

	//Calcular valores dos itens extras
	// extraItems := calcbody.GetExtraItemsMaterials(materials)

	//Não processar, caso não tenha nenhum item extra
	// if len(extraItems) > 0 {

	// 	recipes := repos.GetRecipesForExtra()

	// 	if err == nil {
	// 		//Channel para receber as informações de cálculo de dotz extras
	// 		chanExtraItemResults := make(chan *viewModels.CalculateResultExtraItem)

	// 		for _, extra := range extraItems {
	// 			go calculateExtraItem(chanExtraItemResults, recipes, extra, isPreview)
	// 		}

	// 		for i := 0; i < len(extraItems); i++ {
	// 			select {
	// 			case extraResult := <-chanExtraItemResults:
	// 				if extraResult != nil {
	// 					log.Printf("resultado da goroutine %v", extraResult)
	// 					resultModel.ExtraItems = append(resultModel.ExtraItems, extraResult)
	// 				}
	// 			}
	// 		}
	// 	} else {
	// 		log.Print("Não foi possível obter as fórmulas de itens extras", err)
	// 	}
	// }

	return &resultModel, recipe, nil
}
