package taskThreads

import (
	"fmt"
	"sort"
	"sync"
)

// 2. **Запуск нескольких горутин и ожидание их завершения**

//     > Задача: Напишите программу, которая запускает 5 горутин, каждая из которых печатает свой номер (от 1 до 5),
// и использует sync.WaitGroup для их синхронизации(нужно подождать их выполнения).

//     Можно ли решить задачу без waitGroup? Какие есть варианты?
//     Можно ли сделать так чтобы номера выводились в определённом порядке? Почему? Может всё-таки можно?

//     Как влияет GOMAXPROCS на выполнение программы?
//     >

func TaskTwo() {

	// РЕШЕНИЕ ЗАДАЧИ С WAIT.GROUP
	// Buf для вывода данных в нужном порядке
	buf := make([]int, 0, 5)
	// Кол-во потоков
	threads := 5

	wg := &sync.WaitGroup{}
	ch := make(chan int)

	// Заполнение горутин
	for i := range threads {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- i
		}()
	}

	// Вывод горутин
	go func() {
		for i := range threads {
			buf = append(buf, i)
		}

		// Сортируем массив
		sort.Ints(buf)

		// Выводим в нужном порядке и закрываем горутины
		for i := range ch {
			fmt.Println(buf[i])
		}

	}()

	wg.Wait()

	// РЕШЕНИЕ ЗАДАЧИ БЕЗ WAIT.GROUP
	// threadsCount := 5
	// //Используем горутину с буфером для параллельности работы
	// threads := make(chan int, threadsCount)
	// done := make(chan struct{})

	// for i := 0; i < threadsCount; i++ {
	// 	go func(n int) {
	// 		threads <- n
	// 		done <- struct{}{}
	// 	}(i)
	// }

	// // Анализируем и потом закрываем поток threads
	// go func() {
	// 	for i := 0; i < threadsCount; i++ {
	// 		<-done
	// 	}
	// 	close(threads)
	// }()

	// // Читаем из threads
	// for i := 0; i < threadsCount; i++ {
	// 	ii := <-threads
	// 	fmt.Println(ii)
	// }

}
