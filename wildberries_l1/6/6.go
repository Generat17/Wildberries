/*

6. Реализовать все возможные способы остановки выполнения горутины.

*/

package main

import (
	"context"
	"sync"
	"time"
)

// solve1 естественное завершение горутины
func solve1() {
	// Объявляем WaitGroup
	var wg sync.WaitGroup

	// Добавляем задачу в WaitGroup
	wg.Add(1)

	// Тестовая горутина
	go func(wg *sync.WaitGroup) {
		// Перед заверешением функции отработает счетчик WaitGroup уменьшится
		defer wg.Done()

		// Задержка 1 секунда
		time.Sleep(1 * time.Second)

		return
	}(&wg)

	// Ожидать завершения горутины
	wg.Wait()

	// Если бы мы не использовали WaitGroup то горутина Main могла бы завершиться раньше, чем тестовая
}

// solve2 завершение горутины через сигнал из канала
func solve2() {
	// Создаем канал для отправки стоп-сигнала
	stopCh := make(chan struct{})

	// Объявляем WaitGroup
	var wg sync.WaitGroup

	// Добавляем задачу в WaitGroup
	wg.Add(1)

	// Тестовая горутина
	go func(stopCh chan struct{}, wg *sync.WaitGroup) {
		defer wg.Done()

		// Логика действий в горутине

		// Проверять канал на наличие сигнала остановки
		select {
		case <-stopCh:
			// Получен сигнал остановки, завершить выполнение горутины
			return
		default:
			// Продолжить выполнение действий
		}
	}(stopCh, &wg)

	// Через некоторое время отправить сигнал остановки
	go func() {
		time.Sleep(2 * time.Second)
		close(stopCh)
	}()

	// Ожидать завершения горутины
	wg.Wait()
}

// solve3 завершение горутины с помощью контекста WithCancel
func solve3() {
	// Создаем контекст WithCancel
	ctx, cancel := context.WithCancel(context.Background())

	// Объявляем WaitGroup
	var wg sync.WaitGroup

	// Добавляем задачу в WaitGroup
	wg.Add(1)

	// Тестовая горутина
	go func(ctx context.Context, wg *sync.WaitGroup) {
		defer wg.Done()

		// Логика действий в горутине

		// Проверять контекст на наличие сигнала остановки
		select {
		case <-ctx.Done():
			// Получен сигнал остановки, завершить выполнение горутины
			return
		default:
			// Продолжить выполнение действий
		}
	}(ctx, &wg)

	// Через некоторое время вызвать cancel для остановки горутины
	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()

	// Ожидать завершения горутины
	wg.Wait()
}

// solve4 завершение горутины с помощью контекста WithDeadline или WithTimeout
func solve4() {
	// Создаем контекст WithDeadline
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(4*time.Second))

	// Контекст WithTimeout похож на WithDeadline, отличие в том, что
	// WithTimeout - принимает в качестве второго аргумента время, в которое нужно завершить горутину
	// WithDeadline - принимает в качестве второго аргумента длительность (задержку) через которое нужно завершить горутину
	// ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)

	// Объявляем WaitGroup
	var wg sync.WaitGroup

	// Добавляем задачу в WaitGroup
	wg.Add(1)

	// Тестовая горутина
	go func(ctx context.Context, wg *sync.WaitGroup) {
		defer wg.Done()

		// Логика действий в горутине

		// Проверять контекст на наличие сигнала остановки
		select {
		case <-ctx.Done():
			// Получен сигнал остановки, завершить выполнение горутины
			return
		default:
			// Продолжить выполнение действий
		}
	}(ctx, &wg)

	// Через некоторое время вызвать cancel для остановки горутины
	// В данном случае, дедлайн наступит раньше
	go func() {
		time.Sleep(10 * time.Second)
		cancel()
	}()

	// Ожидать завершения горутины
	wg.Wait()
}

func main() {
	solve1()
	solve2()
	solve3()
	solve4()
}
