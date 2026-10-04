package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	var wg sync.WaitGroup
	// spawn 100 goroutines
	// each does 100 atomic.AddInt64(&counter, 1)
	// wg.Wait()

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 100)
		}()
	}

	wg.Wait()
	fmt.Println(counter)

}
