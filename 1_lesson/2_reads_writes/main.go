package main

import (
	"fmt"
	"sync"
)

func main() {
	storage := make(map[int]int, 1000)

	wg := sync.WaitGroup{}
	ops := 1000
	mu := sync.RWMutex{}

	wg.Add(ops)
	for i := 0; i < ops; i++ {
		i := i
		go func() {
			defer wg.Done()

			mu.Lock()
			defer mu.Unlock()
			storage[i] = i
		}()
	}

	wg.Add(ops)
	for i := 0; i < ops; i++ {
		i := i
		go func() {
			defer wg.Done()

			mu.RLock()
			defer mu.RUnlock()
			_, _ = storage[i]
		}()
	}

	wg.Wait()
	fmt.Println(storage)
}
