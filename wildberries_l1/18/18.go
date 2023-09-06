/*

18.	Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.

*/

package main

import (
	"fmt"
	"sync"
)

// Counter - структура-счетчик, которая хранит значение типа int
type Counter struct {
	val   int
	mutex sync.Mutex
}

// Inc - метод, который инкрементирует значение счетчика на 1 атомарно
func (c *Counter) Inc() {
	//atomic.AddInt(&c.val, 1)
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.val++
}

// Value - метод, который возвращает текущее значение счетчика
func (c *Counter) Value() int {
	//return atomic.LoadInt(&c.val)
	return c.val
}

func main() {
	// Создаем новый счетчик
	counter := &Counter{}

	// Создаем WaitGroup для синхронизации горутин
	var wg sync.WaitGroup

	// Запускаем 1000 горутин, каждая из которых увеличивает счетчик на 1
	for i := 0; i < 1000; i++ {
		wg.Add(1) // Увеличиваем счетчик WaitGroup на 1
		go func() {
			defer wg.Done() // Уменьшаем счетчик WaitGroup на 1 по завершению горутины
			counter.Inc()   // Инкрементируем счетчик
		}()
	}

	wg.Wait() // Ждем, пока все горутины завершатся

	// Выводим итоговое значение счетчика
	fmt.Println(counter.Value())
}
