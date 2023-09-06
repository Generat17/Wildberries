/*

9. Разработать конвейер чисел.
Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout.

*/

package main

import (
	"fmt"
	"sync"
)

// Функция, которая читает числа из массива и пишет их в канал out
func source(nums []int, out chan<- int) {
	for _, n := range nums {
		out <- n
	}
	close(out)
}

// Функция, которая читает числа из канала in, умножает их на 2 и пишет в канал out
func multiply(in <-chan int, out chan<- int) {
	for n := range in {
		out <- n * 2
	}
	close(out)
}

// Функция, которая читает числа из канала in и выводит их в stdout
func sink(in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Сообщаем WaitGroup, что горутина завершилась
	for n := range in {
		fmt.Println(n)
	}
}

func main() {
	// Создаем два канала для передачи чисел между функциями
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Создаем массив чисел для обработки
	nums := []int{1, 2, 3, 4, 5}

	// Создаем WaitGroup для синхронизации горутин
	wg := &sync.WaitGroup{}
	wg.Add(1) // Увеличиваем счетчик WaitGroup на 1

	// Запускаем три горутины для выполнения трех стадий конвейера
	go source(nums, ch1)
	go multiply(ch1, ch2)
	go sink(ch2, wg)

	// Ждем завершения всех горутин
	wg.Wait() // Блокируем main до тех пор, пока счетчик WaitGroup не станет 0
}
