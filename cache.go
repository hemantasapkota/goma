package goma

import (
	"sync"
)

type AppCache struct {
	sync.Mutex
	Objects map[string]DBObject
}

func (c AppCache) Key() string {
	return "goma.cache"
}

var cache *AppCache = &AppCache{Objects: make(map[string]DBObject)}

func GetAppCache() *AppCache {
	return cache
}

func (c *AppCache) Put(obj DBObject) *AppCache {
	c.Lock()
	defer c.Unlock()

	c.Objects[obj.Key()] = obj
	return c
}

func (c *AppCache) Get(obj DBObject) DBObject {
	c.Lock()
	defer c.Unlock()

	o := c.Objects[obj.Key()]
	if o != nil {
		return o
	}
	return obj
}

func (c *AppCache) Delete(object DBObject) {
	c.Lock()
	defer c.Unlock()
	delete(c.Objects, object.Key())
}
