// Package inmemrepo - Пакет для работы с in-memory репозиторием
package inmemrepo

import (
	"github.com/Generat17/dev11/model"
	"sync"
	"time"
)

// InMemoryEventRepo - in-memory реализация репозитория событий
type InMemoryEventRepo struct {
	events []model.Event
	genid  model.EventID
	mutex  sync.Mutex
}

var _ model.EventRepo = &InMemoryEventRepo{}

// NewInMemoryEventRepo - создает in-memory репозиторий событий
func NewInMemoryEventRepo() *InMemoryEventRepo {
	return &InMemoryEventRepo{}
}

// Create implements model.EventRepo.
func (er *InMemoryEventRepo) Create(event model.Event) (model.Event, error) {
	er.mutex.Lock()
	defer er.mutex.Unlock()

	er.genid++
	event.ID = er.genid
	er.events = append(er.events, event)
	return event, nil
}

// Delete implements model.EventRepo.
func (er *InMemoryEventRepo) Delete(id model.EventID) error {
	er.mutex.Lock()
	defer er.mutex.Unlock()

	for i, e := range er.events {
		if e.ID == id {
			er.events = append(er.events[:i], er.events[i+1:]...)
			return nil
		}
	}
	return nil
}

// GetByID implements model.EventRepo.
func (er *InMemoryEventRepo) GetByID(id model.EventID) (model.Event, error) {
	er.mutex.Lock()
	defer er.mutex.Unlock()

	for _, e := range er.events {
		if e.ID == id {
			return e, nil
		}
	}
	return model.Event{}, model.ErrEventNotFound
}

// GetByPeriod implements model.EventRepo.
func (er *InMemoryEventRepo) GetByPeriod(start time.Time, end time.Time) ([]model.Event, error) {
	if start.After(end) {
		return nil, model.ErrInvalidPeriod
	}

	er.mutex.Lock()
	defer er.mutex.Unlock()

	res := make([]model.Event, 0)
	for _, e := range er.events {
		if !e.Date.Before(start) && !e.Date.After(end) {
			res = append(res, e)
		}
	}
	return res, nil
}

// Update implements model.EventRepo.
func (er *InMemoryEventRepo) Update(event model.Event) (model.Event, error) {
	er.mutex.Lock()
	defer er.mutex.Unlock()

	for i, e := range er.events {
		if e.ID == event.ID {
			er.events[i] = event
			return event, nil
		}
	}
	return model.Event{}, model.ErrEventNotFound
}
