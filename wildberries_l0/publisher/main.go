package main

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"os"
	"publisher/model"
	"strconv"
	"time"
)

// конфиг для подключения к NATS streaming
const (
	NATSStreamingURL = "localhost:53616"
	ClusterID        = "test-cluster"
	ClientID         = "test-publisher"
	Channel          = "testch"
)

// задержка между отправкой сообщений
const duration = 5 * time.Second

func main() {

	// читаем тестовый файл
	file, err := os.ReadFile("testData.json")
	if err != nil {
		log.Fatal(err)
	}

	var order model.Order

	// парсим данные из файла и сохраняем в объект
	err = json.Unmarshal(file, &order)
	if err != nil {
		log.Fatal(err)
	}

	// подключаемся к nats streaming серверу
	sc, err := stan.Connect(ClusterID, ClientID+"1", stan.NatsURL(NATSStreamingURL))
	if err != nil {
		log.Fatal(err)
	}

	// бесконечный цикл, с периодичность duration пишет сообщения в канал
	// duration - константа, можно настроить в начале файла
	for {

		// в id заказа записываем текущее unix время, чтобы id-ник всегда был уникальным
		order.OrderUID = strconv.Itoa(int(time.Now().Unix()))

		// из экземпляра структуры делаем json объект
		json, err := json.Marshal(order)

		// пишем json объект в канал
		err = sc.Publish(Channel, json)
		if err != nil {
			log.Panic(err)
		}

		// уведомление об успешной отправке сообщения
		fmt.Println("OK. Message send success with id =", order.OrderUID)

		// задержка между отправкой сообщений
		time.Sleep(duration)
	}
}
