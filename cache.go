package goma

import ()

type AppCache struct {
	Objects map[string]DBObject
}

func (c AppCache) Key() string {
	return "goma.cache"
}

var cache *AppCache

func GetAppCache() *AppCache {
	if cache == nil {
		cache = &AppCache{
			Objects: make(map[string]DBObject),
		}
	}
	return cache
}

func (c *AppCache) Put(obj DBObject) *AppCache {
	c.Objects[obj.Key()] = obj
	return c
}

func (c *AppCache) Get(obj DBObject) DBObject {
	o := c.Objects[obj.Key()]
	if o != nil {
		return o
	}
	return obj
}

func (c *AppCache) Delete(object DBObject) {
	delete(c.Objects, object.Key())
}
