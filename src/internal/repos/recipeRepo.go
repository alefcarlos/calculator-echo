package repos

import (
	"github.com/alefcarlos/calculator-echo/src/internal/caches"
	"github.com/alefcarlos/calculator-echo/src/internal/entities"
	"github.com/patrickmn/go-cache"
)

var recipesCache = caches.NewInMemoryCache(cache.NoExpiration, cache.DefaultExpiration)

//AddRecipe -
func AddRecipe(recipe *entities.Recipe) {
	recipe.ID = recipesCache.ItemsCount() + 1

	recipesCache.AddItem(string(recipe.ID), recipe)
}

//GetRecipes -
func GetRecipes() (recipes []*entities.Recipe) {
	v := make([]*entities.Recipe, 0, 0)

	for _, value := range recipesCache.Items() {
		v = append(v, value.Object.(*entities.Recipe))
	}

	recipes = v
	return
}

// FilterRecipe returns a new slice holding only
// the elements of s that satisfy fn()
func FilterRecipe(s []*entities.Recipe, fn func(*entities.Recipe) bool) []*entities.Recipe {
	var p []*entities.Recipe // == nil
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}

//FindRecipe obtém um item num array de acordo com a função fn()
func FindRecipe(s []*entities.Recipe, fn func(*entities.Recipe) bool) *entities.Recipe {
	for _, v := range s {
		if fn(v) {
			return v
		}
	}

	return nil
}

//IndexRecipe retorna o index do item de acordo com fn()
func IndexRecipe(s []*entities.Recipe, fn func(*entities.Recipe) bool) int {
	index := -1

	for i, v := range s {
		if fn(v) {
			index = i
			break
		}
	}

	return index
}
