package utils

import (
	"sync"
	"time"
)

type TTLMap[K comparable, V any] struct {
	items sync.Map
}

type item[V any] struct {
	value      V
	expireTime time.Time
}

func NewTTLMap[K comparable, V any]() *TTLMap[K, V] {
	m := &TTLMap[K, V]{}
	go m.cleanup()
	return m
}

func (m *TTLMap[K, V]) Set(key K, value V, ttl time.Duration) {
	m.items.Store(key, &item[V]{
		value:      value,
		expireTime: time.Now().Add(ttl),
	})
}

func (m *TTLMap[K, V]) Get(key K) (v V, exists bool) {
	value, ok := m.items.Load(key)
	if !ok {
		return v, false
	}

	item := value.(*item[V])
	if time.Now().After(item.expireTime) {
		m.items.Delete(key)
		return v, false
	}
	return item.value, true
}

func (m *TTLMap[K, V]) Delete(key K) {
	m.items.Delete(key)
}

func (m *TTLMap[K, V]) GetOrSet(key K, defaultValue V, ttl time.Duration) V {
	for {
		if value, exists := m.Get(key); exists {
			return value
		}

		thisItem := &item[V]{
			value:      defaultValue,
			expireTime: time.Now().Add(ttl),
		}

		actual, loaded := m.items.LoadOrStore(key, thisItem)
		if !loaded {
			return defaultValue
		}

		existingItem := actual.(*item[V])
		if !time.Now().After(existingItem.expireTime) {
			return existingItem.value
		}

		m.items.Delete(key)
	}
}

func (m *TTLMap[K, V]) cleanup() {
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		m.items.Range(func(key, value interface{}) bool {
			item := value.(*item[V])
			if time.Now().After(item.expireTime) {
				m.items.Delete(key)
			}
			return true
		})
	}
}

func (m *TTLMap[K, V]) Range(f func(key K, value V) bool) {
	m.items.Range(func(k, v interface{}) bool {
		key := k.(K)
		item := v.(*item[V])

		if time.Now().After(item.expireTime) {
			m.items.Delete(key)
			return true
		}

		return f(key, item.value)
	})
}
