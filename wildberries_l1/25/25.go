/*

25. Реализовать собственную функцию sleep.

*/

package main

import (
	"fmt"
	"time"
)

func mySleep(d time.Duration, otherChannel chan string) {
	timer := time.NewTimer(d) // создать таймер на d
	select {
	case <-timer.C: // ждать, пока таймер не закончится
		return
	case <-otherChannel: // ждать другого события
		timer.Stop() // остановить таймер
		return
	}
}

func main() {
	someOtherChannel := make(chan string) // создать канал для строк

	go func() {
		time.Sleep(3 * time.Second) // спать 3 секунды
		someOtherChannel <- "stop"  // отправить "stop" в канал
	}()

	fmt.Println("Start sleeping")
	mySleep(10*time.Second, someOtherChannel) // вызвать функцию mySleep на 5 секунд
	fmt.Println("End sleeping")
}
