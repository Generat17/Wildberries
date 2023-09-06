## 7. Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Функция принимает переменное кол-во интов, кидает их в канал, закрывает его и возвращает канал
func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

// Функция слияния двух каналов
func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			// Нужна проверка на закрытие канала: v, ok := <- a; 
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
		// Тут нужно закрыть канал
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}
```

## Вывод:

```
1
2
3
4
5
6
7
8
0
...
```

## Ответ:

Программа уходит в бесконечный цикл, потому что в функции merge() нет проверки на то, закрыты каналы или нет. А поскольку из закрытого канала можно читать, функция считывает из каналов нулевые значения и продолжает работать.
