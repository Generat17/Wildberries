/*

12. Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

*/

package main

import "fmt"

// newSet создает новое множество
func newSet(arr []string) map[string]bool {
	set := make(map[string]bool, len(arr))

	// проходим по массиву строк и добавляем каждый элемент в map
	for _, value := range arr {
		set[value] = true
	}

	return set
}

func main() {
	// изначальные данные
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	// создаем множество
	set := newSet(arr)

	fmt.Println(set)
}
