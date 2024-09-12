package utils

import (
	"reflect"
	"sync"
)

var cache sync.Map

func Singleton[T any]() (t *T) {
	hash := reflect.TypeOf(t)
	obj, ok := cache.Load(hash)

	if ok {
		return obj.(*T)
	} else {
		val, _ := cache.LoadOrStore(hash, new(T))
		return val.(*T)
	}
}

func RegisterNewSingleton[T any](t *T, key string) {
	cache.Store(key, t)
}
