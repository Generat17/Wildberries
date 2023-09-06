package main

import (
	"log"
	"solve/http_server"
	"solve/producer"
)

func CreatePostgressConfig() producer.PostgresConfig {
	return producer.PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "root",
		DBName:   "wildberries",
	}
}

func main() {

	// создаем экземпляр структуры, в которой будет хранится кэш
	memCache := producer.NewMemoryCache()

	// Подключение к бд
	client, err := producer.NewPostgresClient(CreatePostgressConfig())
	if err != nil {
		log.Println("Failed connect to postgres instance")
	}

	/*
		в ТЗ указано: "В случае падения сервиса восстанавливать Кеш из Postgres"
		я это понимаю так, что при запуске сервиса, нужно записывать все данные из бд в кэш
		и поэтому сделал именно так.

		Но вообще, я бы использовал другой подход
		1) При поднятии сервиса - кэш пустой;
		-- Пользователь хочет получить запись с id = 1
		2) Смотрим кэш, если нашли запись, то возвращаем ее пользователю;
		3) Если в кэше запись не нашли, смотрим в БД, если запись нашли в БД,
		тогда возвращаем её пользователю и записываем в кэш;
		4) Если запись нигде не нашли, возвращаем ошибку, что запись с таким id не найдена.

		То есть при первой попытке получить запись c id = 1, мы обращаемся к БД,
		а уже при последующих попытках получить ту же самую запись, мы возвращаем ее из кэша.
		Таким образом мы храним в кэше только актуальные заказы.
	*/

	// читаем данные из бд
	orders, err := client.GetOrdersFromPostgres()
	if err != nil {
		log.Println(err)
	}

	// полученные из бд данные, записываем в кэш
	for _, order := range orders {
		memCache.Add(order)
	}

	// подключение к nats steaming
	sc := producer.NewNatsStream()
	sc.RunNatsSteaming(client, memCache)

	// поднятие http сервера
	http_server.StartHTTPServer(client, memCache)
}
