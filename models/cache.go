package models

import (
	"sync"
	"time"
)

// Estructura básica de caché
type SimpleCache struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

// Nueva instancia de caché
func NewSimpleCache() *SimpleCache {
	return &SimpleCache{data: make(map[string]interface{})}
}

// Guardar un valor en caché con TTL
func (c *SimpleCache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	c.data[key] = value
	c.mu.Unlock()

	// Goroutine para eliminar el valor después del TTL
	go func() {
		time.Sleep(ttl)
		c.mu.Lock()
		delete(c.data, key)
		c.mu.Unlock()
	}()
}

// Obtener un valor de la caché
func (c *SimpleCache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, found := c.data[key]
	return value, found
}
