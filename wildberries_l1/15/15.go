/*
15.	К каким негативным последствиям может привести данный фрагмент кода,
и как это исправить? Приведите корректный пример реализации.

var justString string

func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}

----- ОТВЕТ -----

1)	Лучше создать локальную переменную justString в main и передавать ее в качестве аргумента
	- Глобальные переменные храняться в куче, а локальные в стэке
	- Работа кода более предсказуемая

2)	Допустим у нас строка в юникоде, например str := "👋🌎🐶"
	justString = str[:4], justString будет содержать только "👋"
	Происходит это потому что операция взятия среза возвращает n байт, а не n символов
	А один символ unicode может весить до 4 байтов
	Для работы с юникодом нужно использовать слайс rune

3)	Для корректной работы с памятью и gc следует использовать функцию copy

*/

package main

import "fmt"

var justString string

func someFunc() {
	v1 := []rune(createHugeString(1 << 10))
	v2 := make([]rune, 100)

	// копируем данные в новую область памяти
	copy(v2, v1[:100])

	justString = string(v2)

	fmt.Println(justString)
}

func createHugeString(length int) string {
	s := make([]rune, length)

	for i := range s {
		s[i] = '👋'
	}

	//fmt.Println(string(s))

	return string(s)
}

func main() {
	someFunc()
}
