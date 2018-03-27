package caches

import (
	"time"

	"github.com/patrickmn/go-cache"
)

//InMemoryCache estrutura para reutilização de cache
type InMemoryCache struct {
	cache *cache.Cache

	//DefaultExpiration tempo padrão de expiração dos itens do cache
	DefaultExpiration time.Duration

	//DefaultCleanup tempo padrão para limpeza dos itens expirados
	DefaultCleanup time.Duration
}

//AddItem adiciona um item ao cache
func (cache *InMemoryCache) AddItem(ID string, item interface{}) {
	cache.cache.Set(ID, item, cache.DefaultExpiration)
}

//Clear elimna todos os itens do cache
func (cache *InMemoryCache) Clear() {
	cache.cache.Flush()
}

//Items Obtém todos os itens do cache
func (cache *InMemoryCache) Items() map[string]cache.Item {
	return cache.cache.Items()
}

//ItemsCount retorna a quantidade de itens no cache
func (cache *InMemoryCache) ItemsCount() int {
	return cache.cache.ItemCount()
}

//NewInMemoryCache inicializa e retorna uma nova estrutura de cache
//defaultExpiration tempo padrão de expiração, se for <= 1, nunca irá expirar
//defaultCleanup tempo padrão para limpar os itens expirados, se for <= 1, nunca irá remover os itens expirados
func NewInMemoryCache(defaultExpiration, defaultCleanup time.Duration) *InMemoryCache {
	return &InMemoryCache{
		DefaultExpiration: defaultExpiration,
		DefaultCleanup:    defaultCleanup,
		cache:             cache.New(defaultExpiration, defaultCleanup),
	}
}
