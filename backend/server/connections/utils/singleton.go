package utils

import (
	"sync"
)

var cache sync.Map

// Registra un nuovo singleton.
// Un oggetto viene passato (t) e una chiave unica (key).
// L'oggetto viene registrato e può essere richiamato a bisogno.
//
// Se ritorna false, il risultato non è stato registrato,
// in quanto la chiave è duplicata. Altrimenti, ritorna true.
func RegisterNewSingleton[T any](t *T, key string) bool {
	if val, err := cache.Load(key); err || val != nil {
		return false
	} else {
		cache.Store(key, t)
		return true
	}
}
