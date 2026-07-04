package runtimememorymodel

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Add(delta int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value += delta
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}
