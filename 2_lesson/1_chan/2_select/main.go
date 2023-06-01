// задача про тайм аут
package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type RPC struct {
	jobs chan resp
	rand *rand.Rand
}

type resp struct {
	id  int
	err error
}

func (r *RPC) Call(ctx context.Context) { // передаём контекст, чтобы загасить ожидание
	select {
	case <-ctx.Done(): //выполнился контекст
		r.jobs <- resp{
			id:  0,
			err: errors.New("timeout expired"), // таймаут вышел
		}
	case <-time.After(time.Second * time.Duration(r.rand.Intn(5))): // sleep 0-4 sec
		r.jobs <- resp{
			id: rand.Int(),
		}
	}
}

func main() {
	rpc := &RPC{
		jobs: make(chan resp),
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	ctx := context.Background()                      //фича для отмены затянувшейся горутины
	ctx, _ = context.WithTimeout(ctx, time.Second*2) // вторая переменная cansel

	//обязательно нужно отменить после заданного тайм аута, в противном случае будет отъедать память
	go rpc.Call(ctx) //вызов удалённой процедуры

	//в примере ниже мы заблокировались, других горутин нет и нет команды cansel()
	// <-ctx.Done() //метод контекста, возвращающий канал из которого мы что-то прочтём, если контекст отменили, а если нет
	// то ничего не сможем прочитать и залочимся

	//cansel()

	res := <-rpc.jobs

	fmt.Println(res.id, res.err)
}
