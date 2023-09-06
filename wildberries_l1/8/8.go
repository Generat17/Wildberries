/*

8. Дана переменная int64.
Разработать программу которая устанавливает i-й бит в 1 или 0.

*/

package main

import (
	"fmt"
)

// bitToOne устанавливает i-й бит в 1
func bitToOne(n int64, pos int) int64 {
	n |= 1 << pos
	return n
}

// bitToZero устанавливает i-й бит в 0
func bitToZero(n int64, pos int) int64 {
	n &^= 1 << pos
	return n
}

func main() {

	// Создаем переменную типа int64 для хранения числа
	var num int64
	// Создаем переменную типа int для хранения позиции бита, который мы хотим изменить
	var pos int
	// Создаем переменную типа bool для определения нужной операции над битом
	var value bool

	fmt.Print("Введите число: ")
	// Считываем значение переменной num
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Некорректный ввод")
		return
	}

	fmt.Print("Введите номер бита, который нужно изменить: ")
	// Считываем значение переменной pos
	_, err = fmt.Scan(&pos)
	if err != nil {
		fmt.Println("Некорректный ввод")
		return
	}

	fmt.Print("Введите новое значение для бита (0 или 1): ")
	// Считываем значение переменной value
	_, err = fmt.Scan(&value)
	if err != nil {
		fmt.Println("Некорректный ввод")
		return
	}

	// Выводим ее значение в двоичном формате
	fmt.Printf("%064b\n", num)

	if value {
		num = bitToOne(num, pos)
	} else {
		num = bitToZero(num, pos)
	}

	// Выводим новое значение переменной в двоичном формате
	fmt.Printf("%064b\n", num)
	// Выводим новое значение переменной в десятичном формате
	fmt.Println(num)
}
