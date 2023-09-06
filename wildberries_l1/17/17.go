/*

17. Реализовать бинарный поиск встроенными методами языка.

*/

package main

import "fmt"

// краткое описание алгоритма "Бинарный поиск" (логарифмическое время)
// на вход поступает отсортированный массив, нужно найти в нем элемент и вернуть его индекс
// для этого находим элемент по середине массива и сравниваем с искомым числом,
// если сред.элем больше нашего числа, тогда убираем правую часть
// если меньше, убираем левую и проделываем это, до тех пор, пока не найдем наше число

// return index element and count steps
func binarySearch(sortedList []int, target int) (int, int, error) {
	leftIdx := 0
	rightIdx := len(sortedList) - 1
	countStep := 0

	for leftIdx <= rightIdx {
		countStep++

		mid := (leftIdx + rightIdx) / 2

		guess := sortedList[mid]

		switch {

		case guess == target:
			return mid, countStep, nil

		case guess < target:
			leftIdx = mid + 1

		case guess > target:
			rightIdx = mid - 1

		}
	}

	return -1, 0, fmt.Errorf("target not found")
}

func main() {
	arr := []int{1, 12, 42, 65, 76, 87, 95, 99}

	fmt.Println(binarySearch(arr, 42))
	fmt.Println(binarySearch(arr, 43))
}
