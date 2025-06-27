package taskThreads

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"sync"
	"time"
)

// 6. Контексты

//     > 📦 ЗАДАЧА: Батчевая обработка
//     >
//     >
//     > **Условие**:
//     >
//     > Реализуй функцию `StartBatchProcessor(ctx context.Context, input <-chan int)`, которая:
//     >
//     > - Собирает числа из канала `input` в батчи по максимум 5 элементов.
//     > - Если в течение 2 секунд батч не собран — обрабатывает то, что есть.
//     > - Обработка батча — это просто `fmt.Println("Processed batch:", batch)`.
//     > - Выход из функции должен происходить при отмене контекста (`ctx.Done()`).
//     >
//     > **Дополнительно**:
//     >
//     > - Отмена должна происходить либо через `context.WithTimeout`, либо вручную через `cancel()` — попробовать оба варианта
//     >
//     >  Начальный код с вызовом(доработать)

func TaskSix() {

	allCount := 100

	wg := &sync.WaitGroup{}
	// инициализация канала
	input := make(chan int)
	/* создание контекста  */
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	go func() {
		fmt.Println("Нажмите Enter для отмены")
		bufio.NewReader(os.Stdin).ReadString('\n')
		cancel()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i <= allCount; i++ {
			time.Sleep(time.Millisecond * 25)
			select {
			case input <- i * 2:
			case <-ctx.Done():
			}
		}
		close(input)
	}()

	wg.Add(1)
	go startBatchProcessor(ctx, input, wg)

	wg.Wait()
	fmt.Println("Main: processing stopped")
}

func startBatchProcessor(ctx context.Context, input <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	batchSize := 5
	batchSlice := make([]int, 0, batchSize)

	// Читаем данные
	for {
		select {
		// -- По требованию
		case <-ctx.Done():
			return
		// -- Полное успешное выполнение пакета
		case i, ok := <-input:

			if !ok {
				return
			}
			if len(batchSlice) == batchSize {

				fmt.Println("Processed batch:", batchSlice)
				batchSlice = batchSlice[:0]
			}
			batchSlice = append(batchSlice, i)
		// -- После 2х секунд
		case <-time.After(time.Second * 2):
			if len(batchSlice) > 0 {
				fmt.Println("Processed batch:", batchSlice)
			}
			return
		}

	}
}
