package ttlcache

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestListCache(t *testing.T) {
	lc := NewListCache(5*time.Second, 5, func(key string, list ListData) {
		fmt.Println("cb:", key, "list:", list)
	})

	keys := []string{"a", "b", "c", "d", "e"}
	for i := 0; i < 1000; i++ {
		y := rand.Intn(len(keys))
		lc.Put(keys[y], i)

		//fmt.Println(y, keys[y])

		<-time.After(500 * time.Millisecond)
	}
}
