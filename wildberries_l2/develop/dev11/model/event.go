// Package model - Пакет доменных сущностей и интерфейсов для работы с ними
package model

import (
	"errors"
	"time"
)

// EventID - тип идентификатора события
type EventID int

// Event - сущность события
type Event struct {
	ID          EventID
	Date        time.Time
	Description string
}

var (
	// ErrEventExists - событие уже существует
	ErrEventExists = errors.New("event with given id already exists")
	// ErrEventNotFound - событие не найдено
	ErrEventNotFound = errors.New("event with given id not found")
	// ErrInvalidPeriod - указан неверный временной период
	ErrInvalidPeriod = errors.New("invalid period")
)

// EventRepo - интерфейс репозитория для работы с событиями
type EventRepo interface {
	Create(Event) (Event, error)
	GetByID(EventID) (Event, error)
	GetByPeriod(start time.Time, end time.Time) ([]Event, error)
	Update(Event) (Event, error)
	Delete(EventID) error
}

// EventSerializer - интерфейс сериализации событий
type EventSerializer interface {
	Serialize(Event) ([]byte, error)
	SerializeArray([]Event) ([]byte, error)
}
