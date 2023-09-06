/*

11. Реализовать пересечение двух неупорядоченных множеств.

*/

package main

import (
	"fmt"
)

// Функция, которая возвращает пересечение двух неупорядоченных множеств
func intersection(set1, set2 map[int]bool) map[int]bool {
	// Создаем пустое множество для хранения результата
	result := make(map[int]bool)
	// Проходим по каждому элементу первого множества
	for x := range set1 {
		// Проверяем, есть ли он во втором множестве
		if set2[x] {
			// Если да, то добавляем его в результат
			result[x] = true
		}
	}
	// Возвращаем результат
	return result
}

func main() {

	// Создаем два неупорядоченных множества с некоторыми элементами
	set1 := map[int]bool{1: true, 2: true, 3: true, 4: true}
	set2 := map[int]bool{3: true, 4: true, 5: true, 6: true}

	// Вызываем функцию пересечения и выводим результат
	fmt.Println(intersection(set1, set2))

}
