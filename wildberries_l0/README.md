### команда для запуска nats-streaming сервера в докере
docker run -p 53616:4222 nats-streaming:0.24.6 --cluster_id test-cluster

### скрипт для записи данных в канал - publisher/main.go

### основной сервис main.go
