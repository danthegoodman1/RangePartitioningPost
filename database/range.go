package database

import (
	"strings"
)

type Range struct {
	Low  string
	High string

	KV map[string]string // optimally a btree
}

var (
	// Inf is a fake infinity, just a lot of really high char values, anything above this will break the eval logic
	Inf = strings.Repeat("~", 1000)
)

func (r *Range) OwnsKey(key string) bool {
	return key >= r.Low && key < r.High
}

func (r *Range) Get(key string) *string {
	val, exists := r.KV[key]
	if !exists {
		return nil
	}
	return &val
}

func (r *Range) Delete(key string) {
	delete(r.KV, key)
}

func (r *Range) Set(key, value string) {
	r.KV[key] = value
}
