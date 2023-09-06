/*

13. Поменять местами два числа без создания временной переменной.

*/

package main

import "fmt"

func main() {

	a := 5
	b := 7

	fmt.Println("a =", a, ", b =", b)

	a, b = b, a

	fmt.Println("a =", a, ", b =", b)
}
