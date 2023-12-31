/*

5. Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	// Создаем канал для передачи и приема значений
	ch := make(chan int)

	// Открываем стандартный поток ввода
	in := bufio.NewReader(os.Stdin)

	// Объявляем переменную для хранения времени работы программы в секундах
	var n int

	// Считываем значение переменной n
	_, err := fmt.Fscan(in, &n)
	if err != nil {
		fmt.Println("Некорректный ввод")
		return
	}

	// Запускаем горутину, которая будет отправлять значения в канал
	go func() {
		for i := 0; i < n; i++ {
			ch <- i                 // Отправляем значение i в канал
			time.Sleep(time.Second) // Ждем одну секунду
		}
		close(ch) // Закрываем канал по окончании цикла
	}()

	// Читаем значения из канала в основной горутине
	for v := range ch {
		fmt.Println(v) // Выводим значение на экран
	}
}
