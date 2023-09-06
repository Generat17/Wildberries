package producer

import (
	"errors"
	"sync"
)

/*
	Кэш реализован как хэш-таблица, где ключ - это OrderUID, а значение Order
*/

type MemoryCache struct {
	cache map[string]Order
	mutex sync.Mutex
}

// NewMemoryCache конструктор
func NewMemoryCache() *MemoryCache {

	var mutex sync.Mutex

	return &MemoryCache{
		make(map[string]Order),
		mutex,
	}
}

// Add добавляет новый Order в кэщ
func (m *MemoryCache) Add(order Order) {
	m.mutex.Lock()
	m.cache[order.OrderUID] = order
	m.mutex.Unlock()
}

// Get ищет в кэше Order по его OrderUID
func (m *MemoryCache) Get(id string) (Order, error) {

	defer m.mutex.Unlock()
	m.mutex.Lock()
	order, ok := m.cache[id]

	if !ok {
		return Order{}, errors.New("order not found")
	}

	return order, nil
}
