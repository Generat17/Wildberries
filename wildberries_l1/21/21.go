/*

21.	Реализовать паттерн «адаптер» на любом примере.

*/

package main

import "fmt"

// Target Интерфейс, описывающий целевой интерфейс (тот интерфейс с которым наша система хотела бы работать);
type Target interface {
	Greet()
}

// Adaptee Структура (класс), который наша система должна адаптировать под себя;
type Adaptee struct {
}

// Adapter Структура (класс), адаптер реализующий целевой интерфейс.
type Adapter struct {
	*Adaptee
}

// SayHello метод структуры Adaptee
func (a *Adaptee) SayHello() {
	fmt.Println("Hello world!")
}

// Greet метод адаптера
func (a *Adapter) Greet() {
	a.SayHello()
}

// NewAdapter конструктор адаптера
func NewAdapter(adaptee *Adaptee) Target {
	return &Adapter{adaptee}
}

// функция, принимающая в виде аргумент в виде объекта, реализующего интерфейс Target
func check(target Target) {
	target.Greet()
}

func main() {
	var adaptee Adaptee

	// с помощью адаптера преобразуем Adaptee в необходимый объект интерфейса Target
	target := NewAdapter(&adaptee)

	check(target)
}
