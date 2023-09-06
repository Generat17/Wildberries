package producer

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"log"
)

const (
	NATSStreamingURL = "localhost:53616"
	ClusterID        = "test-cluster"
	ClientID         = "test-publisher"
	Channel          = "testch"
)

type NatsStream struct {
	client stan.Conn
}

// NewNatsStream подключение к NATS-streaming серверу
func NewNatsStream() *NatsStream {
	sc, err := stan.Connect(
		ClusterID, ClientID,
		stan.NatsURL(NATSStreamingURL),
	)
	if err != nil {
		log.Fatalf("error connection nats, %s", err.Error())
	}
	log.Println("nats connection successful")
	return &NatsStream{client: sc}
}

// RunNatsSteaming подписка на канал и чтение сообщение из него
func (n *NatsStream) RunNatsSteaming(client *PostgresClient, cache *MemoryCache) {
	_, err := n.client.Subscribe(
		Channel, func(m *stan.Msg) {
			var order Order

			// читаем сообщение из канала и записываем в структуру Order
			err := json.Unmarshal(m.Data, &order)
			if err != nil {
				log.Fatal(err)
				return
			}

			// пытаемся вставить значение в БД
			err = client.InsertOrder(order)
			if err != nil {
				log.Fatal(err)
				return
			} else {
				// если вставка в бд не вернула ошибку, тогда добавляем запись в кэш
				cache.Add(order)
			}
		})
	if err != nil {
		log.Fatal(err)
	}
}
