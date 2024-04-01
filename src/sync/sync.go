package sync

type Counter struct {
	count int64
}

func (c *Counter) Inc() {
	c.count++
}
func (c *Counter) Value() (value int64) {
	return c.count
}
