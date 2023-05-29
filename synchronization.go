// разминочная задача на синхронизацию
package main

import (
	"fmt"
	"sync"
)

// код возводит числа в квадрат от 0 до 20
func main() {
	wg := sync.WaitGroup{}
	counter := 20
	for i := 0; i <= counter; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(i * i)
		}()
	}

	wg.Wait()

	//time.Sleep(time.Second)
}
