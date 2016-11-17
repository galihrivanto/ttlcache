package ttlcache

import (
	"time"
)

type ListData []interface{}

type CallbackFunc func(key string, list ListData)

type ListCache struct {
	cache     *Cache
	maxLength int
	cbfunc    CallbackFunc
}

func (lc *ListCache) Put(key string, data interface{}) {
	var list ListData
	old, found := lc.cache.Get(key)
	if found {
		list = old.(ListData)
	}

	// merge new date with old list
	list = append(list, data)

	if len(list) >= lc.maxLength {
		if lc.cbfunc != nil {
			lc.cbfunc(key, list)
		}

		lc.cache.Del(key)
	} else {
		lc.cache.Set(key, list)
	}
}

func NewListCache(d time.Duration, l int, cbfunc CallbackFunc) *ListCache {
	cache := NewCache(d)
	cache.expfunc = func(key string, data interface{}) {
		list, ok := data.(ListData)
		if ok {
			cbfunc(key, list)
		}
	}

	return &ListCache{
		cache:     cache,
		maxLength: l,
		cbfunc:    cbfunc,
	}
}
