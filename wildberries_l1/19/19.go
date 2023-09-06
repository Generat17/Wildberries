/*

19. Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
Символы могут быть unicode.

*/

package main

import (
	"fmt"
)

// Reverse - функция, которая переворачивает строку
func Reverse(s string) string {
	// Преобразуем строку в срез рун
	runes := []rune(s)

	// Обменяем местами элементы среза с начала и с конца
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Преобразуем срез рун обратно в строку
	return string(runes)
}

func main() {
	// Тестируем функцию на разных строках
	fmt.Println(Reverse("главрыба")) // абырвалг
	fmt.Println(Reverse("Hello"))    // olleH
	fmt.Println(Reverse("Привет"))   // тевирП
	fmt.Println(Reverse("👋🌎"))       // 🌎👋
}
