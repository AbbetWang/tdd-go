package sync

import "sync"

type Counter struct {
	mu    sync.Mutex
	count int64
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}
func (c *Counter) Value() (value int64) {
	return c.count
}
