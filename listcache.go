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

// Put add new item to list, return current list length
// if not exceeded max length
func (lc *ListCache) Put(key string, data interface{}) int {
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

		return -1

	} else {
		lc.cache.Set(key, list)
	}

	return len(list)
}

func (lc *ListCache) Get(key string) (ListData, bool) {
	l, found := lc.cache.Get(key)
	if found {
		return l.(ListData), true
	}

	return nil, false
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
