package sync_lib

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	count uint64
}

func (c *Counter) Incr() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func SyncWaitGroup() {
	var c Counter
	var wg sync.WaitGroup
	wg.Add(10)
	for i :=0;i<10;i++{
		go func() {
			defer wg.Done()
			c.Incr()
		}()
	}

	wg.Wait()
	fmt.Println(c.Count())
}
