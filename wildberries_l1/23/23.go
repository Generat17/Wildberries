/*

23. Удалить i-ый элемент из слайса.

*/

package main

import "fmt"

// removeElement удаляет i-ый элемент из слайса
func removeElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	index := 5

	fmt.Println(slice)

	// Удаление i-го элемента
	slice = removeElement(slice, index)

	fmt.Println(slice)
}
