/*

1. Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

*/

package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

type Car struct {
	Model    string
	MaxSpeed int
}

func (h *Human) SayHello() {
	fmt.Println("Hello, my name is", h.Name)
}

type Action struct {
	Human
	Car
}

func (a *Action) Move() {
	fmt.Println(a.Name, "is move")
	fmt.Println(a.Model, "is move")
}

func main() {
	human := Action{
		Human: Human{
			Name: "Alex Smith",
			Age:  25,
		},
	}

	car := Action{
		Car: Car{
			Model:    "BMW",
			MaxSpeed: 200,
		},
	}

	human.Move()
	car.Move()

}
