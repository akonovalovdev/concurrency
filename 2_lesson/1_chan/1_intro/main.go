package main

import (
	"fmt"
	"runtime"
	"time"
)

// ?
//func main_() {
//	ch := make(chan int)                // запись в небуферизованный канал без чтения - дедлок
//	ch := make(chan int, 1)             // буферизованный канал
//	ch <- 1                             // при записи в буферизованный канал ошибки нет, даже без чтения
//	fmt.Println(runtime.NumGoroutine()) //печать количества работающих горутин
//	<-ch                                // чтение без записи из любого канала - дедлок
//}

// запускаем горутину которая в цикле раз в секунду пишет значение i
func main() {
	fmt.Println(runtime.NumGoroutine())
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			fmt.Println("iter: ", i)
		}
		fmt.Println(runtime.NumGoroutine())
	}()
	select {} //пустой селескт - способ себе устроить дедлок)
	fmt.Println("finish")
}
