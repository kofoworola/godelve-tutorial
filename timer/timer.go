package timer

import "time"

type CustomTick struct {
	interval time.Duration
	count    int

	runnable func() error
}

func New(count int, interval time.Duration, runnable func() error) *CustomTick {
	return &CustomTick{
		count:    count,
		interval: interval,
		runnable: runnable,
	}
}

func (c *CustomTick) Begin() chan error {
	respChan := make(chan error)
	go func() {
		for i := 0; i < c.count; i++ {
			if err := c.runnable(); err != nil {
				respChan <- err
				break
			}
			time.Sleep(c.interval)
		}
		respChan <- nil
	}()

	return respChan
}
