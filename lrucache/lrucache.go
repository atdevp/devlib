package lrucache

import (
	"sync"

	"github.com/atdevp/devlib/lrucache/lru"
)

type Cache struct {
	lru  *lru.LRU
	lock sync.RWMutex
}

func New(size int) (*Cache, error) {

	c, e := lru.NewLRU(size)
	if e != nil {
		return nil, e
	}
	return &Cache{lru: c}, nil
}

func (c *Cache) Set(key, value interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.lru.Set(key, value)
}

func (c *Cache) Contains(key interface{}) bool {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.lru.Exist(key)
}

func (c *Cache) Remove(key interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.lru.Delete(key)
}

func (c *Cache) Len() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.lru.Length()
}

func (c *Cache) RemoveOldest() {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.lru.RemoveOldest()
}
