package internal

import (
	"sync"
	"time"
)

type Cache struct {
	data  map[string]cacheEntry
	mutex *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	myCache := Cache{
		data:  map[string]cacheEntry{},
		mutex: &sync.Mutex{},
	}
	go reapLoop(&myCache, interval)
	return myCache
}

func reapLoop(c *Cache, interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		<-ticker.C
		c.mutex.Lock()
		for key, value := range c.data {
			if value.createdAt.Before(time.Now().Add(-1 * interval)) {
				delete(c.data, key)
			}
		}
		c.mutex.Unlock()
	}
}

func (c Cache) Add(k string, v []byte) {
	c.mutex.Lock()
	c.data[k] = cacheEntry{
		createdAt: time.Now(),
		val:       v,
	}
	c.mutex.Unlock()
}

func (c Cache) Get(k string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	entry, ok := c.data[k]
	if !ok {
		return nil, false
	}
	return entry.val, true
}
