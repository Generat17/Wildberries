/*

16. Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Num interface {
	int | int8 | float64
}

func partition[T Num](arr []T, low int, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++

			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// quickSort быстрая сортировка
func quickSort[T Num](arr []T, low int, high int) {
	if low < high {
		pi := partition[T](arr, low, high)

		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

// compArr сравнивает два массива
func compArr(arr []int, n int) bool {
	check := true

	var trueArr = make([]int, n)
	for i := 0; i < n; i++ {
		trueArr[i] = i
	}

	for i := 0; i < n; i++ {
		if arr[i] != trueArr[i] {
			check = false
		}
	}

	// if arrays compare then true, else false
	return check
}

func main() {

	rand.Seed(time.Now().Unix())

	const n = 600  // длинна массива
	const m = 4000 // кол-во повторений

	start := time.Now() // начало измерения

	flag := true
	for i := 0; i < m; i++ {
		randArr := rand.Perm(n)
		quickSort[int](randArr, 0, len(randArr)-1)
		if !compArr(randArr, n) {
			flag = false
		}
	}

	duration := time.Since(start)
	fmt.Println(duration)

	fmt.Println(flag)
}
