package main

// сценарий отработки зависаний при записи и чтении из каналов
import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		select {
		case val := <-ch: // по аналоги со switch, чтобы не словить дедлок(Dead Lock), необходимо указать default
			fmt.Println(val)
		default: //значение по умолчанию
			fmt.Println("Default") // дефолтный сценарий, если другие не выполнились
		}
	}()
	time.Sleep(1 * time.Second)

	fmt.Println(runtime.NumGoroutine()) //счётчик запущеных горутин

}
