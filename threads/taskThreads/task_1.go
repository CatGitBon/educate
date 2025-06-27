package taskThreads

import (
	"context"
	"fmt"
	"time"
)

// 1-я задача
// 1. **Запуск горутины и ожидание её завершения**

//     > Задача: Напишите функцию, которая запускает горутину, выполняющую fmt.Println("Hello from goroutine!"), и использует sync.WaitGroup для ожидания её завершения.

//     Какие способы есть ещё кроме waitGroup, чтобы дождаться выполнения горутины? Приведи хотя бы 2 примера.
//     >

func TaskOne() {

	// wg := &sync.WaitGroup{}

	// wg.Add(1)
	// ch := make(chan int)
	ctx, cansel := context.WithTimeout(context.Background(), time.Second*2)
	defer cansel()

	go explainGorutine()

	<-ctx.Done()
	// wg.Wait()
	// <-ch
}

func explainGorutine() {
	fmt.Println("Hello from goroutine!")
	// close(ch)
	// wg.Done()
}
