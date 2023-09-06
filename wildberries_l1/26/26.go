/*

26. Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.

Например:

abcd — true
abCdefAaf — false
aabcd — false

*/

package main

import (
	"fmt"
	"strings"
)

// Функция, которая проверяет, что все символы в строке уникальные
func isUnique(s string) bool {
	// Приводим строку к нижнему регистру
	s = strings.ToLower(s)
	// Создаем множество для хранения встреченных символов
	set := make(map[rune]bool)
	// Проходим по строке и проверяем, есть ли символ в множестве
	for _, r := range s {
		if set[r] {
			// Если символ уже есть в множестве, значит строка не уникальна
			return false
		}
		// Добавляем символ в множество
		set[r] = true
	}
	// Если мы дошли до конца строки, значит все символы уникальны
	return true
}

func main() {
	// Тестируем функцию на разных примерах
	fmt.Println(isUnique("abcd"))      // true
	fmt.Println(isUnique("abbcd"))     // false
	fmt.Println(isUnique("abCdefAaf")) // false
	fmt.Println(isUnique("  abcd"))    // false
	fmt.Println(isUnique(" abcd"))     // true
}
