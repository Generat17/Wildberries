/*

7. Реализовать конкурентную запись данных в map.

*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем map для хранения данных
	data := make(map[string]string)

	// Создаем мьютекс для синхронизации доступа к map
	var mu sync.Mutex

	// Создаем группу для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Запускаем 10 горутин, которые будут записывать данные в map
	for i := 0; i < 10; i++ {
		wg.Add(1) // Увеличиваем счетчик группы
		go func(i int) {
			defer wg.Done()                    // Уменьшаем счетчик группы по завершении горутины
			key := fmt.Sprintf("key%d", i)     // Формируем ключ
			value := fmt.Sprintf("value%d", i) // Формируем значение
			mu.Lock()                          // Блокируем мьютекс перед записью в map
			data[key] = value                  // Записываем значение в map по ключу
			mu.Unlock()                        // Разблокируем мьютекс после записи в map
		}(i)
	}

	wg.Wait() // Ждем, пока все горутины закончат работу

	fmt.Println("Данные в map:")
	for k, v := range data {
		fmt.Printf("%s: %s\n", k, v)
	}
}
