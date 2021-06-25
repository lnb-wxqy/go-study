package sync_lib

import (
	"fmt"
	"sync"
)

func sync_waitgroup() {
	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}

	wg.Wait()
}
