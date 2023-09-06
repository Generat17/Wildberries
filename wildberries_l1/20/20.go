/*

20. Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».

*/

package main

import (
	"fmt"
	"strings"
)

// ReverseWords - функция, которая переворачивает слова в строке
func ReverseWords(s string) string {
	// Разбиваем строку на срез строк по пробелам
	words := strings.Split(s, " ")

	// Обменяем местами элементы среза с начала и с конца
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Соединяем элементы среза в одну строку с пробелами
	return strings.Join(words, " ")
}

func main() {
	// Тестируем функцию на разных строках
	fmt.Println(ReverseWords("snow dog sun")) // sun dog snow
	fmt.Println(ReverseWords("Hello World"))  // World Hello
	fmt.Println(ReverseWords("Привет Мир"))   // Мир Привет
	fmt.Println(ReverseWords("👋 🌎 🐶"))        // 🐶🌎👋
}
