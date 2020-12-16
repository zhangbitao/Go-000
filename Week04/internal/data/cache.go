package data

import (
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Cacher interface {
	Get(key string) (value string)
	Set(key string, value []string)
}

var _ Cacher = (*cache)(nil)

func NewCache() Cacher {
	return &cache{
		qa: make(map[string][]string),
	}
}

type cache struct {
	mu sync.RWMutex
	qa map[string][]string
}

func (c *cache) Get(key string) (value string) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	values, ok := c.qa[key]
	if !ok {
		return ""
	}
	if len(values) == 1 {
		return values[0]
	}

	return values[rand.Intn(len(values))]
}

func (c *cache) Set(key string, value []string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.qa[key] = append(c.qa[key], value...)
}
