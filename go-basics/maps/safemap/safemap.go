package safemap

import (
	"fmt"
	"sync"
)

type SafeMap[K comparable, V any] struct {
	mu   sync.RWMutex
	data map[K]V
}

// Constructor
func New[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{
		data: make(map[K]V),
	}
}

func (m *SafeMap[K, V]) Insert(key K, value V) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}

func (m *SafeMap[K, V]) Get(key K) (V, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	value, ok := m.data[key]
	if !ok {
		return value, fmt.Errorf("key %v not found", key)
	}
	return value, nil
}

func (m *SafeMap[K, V]) Update(key K, value V) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, ok := m.data[key]
	if !ok {
		return fmt.Errorf("key %v not found", key)
	}
	m.data[key] = value
	return nil
}

func (m *SafeMap[K, V]) Delete(key K) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, ok := m.data[key]
	if !ok {
		return fmt.Errorf("key %v not found", key)
	}

	delete(m.data, key)
	return nil
}

func (m *SafeMap[K, V]) HasKey(key K) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, ok := m.data[key]
	return ok
}
