/*

14. Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.

*/

package main

import (
	"fmt"
)

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("x is an int")
	case string:
		fmt.Println("x is a string")
	case bool:
		fmt.Println("x is a bool")
	case chan int:
		fmt.Println("x is a channel of int")
	default:
		fmt.Println("x is something else")
	}
}

func main() {
	var x interface{}

	x = 20
	checkType(x)

	x = "hello world"
	checkType(x)

	x = true
	checkType(x)

	x = make(chan int)
	checkType(x)
}
