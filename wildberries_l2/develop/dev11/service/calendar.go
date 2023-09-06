// Package service - Пакет с сервисом для работы с событиями
package service

import (
	"errors"
	"github.com/Generat17/dev11/model"
	"time"
)

var (
	// ErrInvalidDateFormat - Не корректная дата
	ErrInvalidDateFormat = errors.New("invalid date format")
)

// EventCreateDTO - DTO для создания события
type EventCreateDTO struct {
	Date        string
	Description string
}

// EventUpdateDTO - DTO для обновления события
type EventUpdateDTO struct {
	ID          model.EventID
	Date        *string
	Description *string
}

// CalendarService - Сервис календаря
type CalendarService struct {
	eventRepo  model.EventRepo
	dateFormat string
}

// NewCalendarService - конструктор
func NewCalendarService(eventRepo model.EventRepo, dateFormat string) *CalendarService {
	return &CalendarService{
		eventRepo:  eventRepo,
		dateFormat: dateFormat,
	}
}

// CreateEvent - Создание события
func (c *CalendarService) CreateEvent(dto EventCreateDTO) (model.Event, error) {
	date, err := time.Parse(c.dateFormat, dto.Date)
	if err != nil {
		return model.Event{}, ErrInvalidDateFormat
	}
	return c.eventRepo.Create(model.Event{
		Date:        date,
		Description: dto.Description,
	})
}

// UpdateEvent - Обновление события
func (c *CalendarService) UpdateEvent(dto EventUpdateDTO) (model.Event, error) {
	event, err := c.eventRepo.GetByID(dto.ID)
	if err != nil {
		return model.Event{}, err
	}
	if dto.Date != nil {
		date, err := time.Parse(c.dateFormat, *dto.Date)
		if err != nil {
			return model.Event{}, ErrInvalidDateFormat
		}
		event.Date = date
	}
	if dto.Description != nil {
		event.Description = *dto.Description
	}
	return c.eventRepo.Update(event)
}

// DeleteEvent - Удаление события
func (c *CalendarService) DeleteEvent(ID model.EventID) error {
	c.eventRepo.Delete(ID)
	return nil
}

// GetForDay - Получение событий за текущий день
func (c *CalendarService) GetForDay() ([]model.Event, error) {
	//TODO: test and fix
	start := time.Now().Add(-12 * time.Hour).Round(24 * time.Hour)
	end := start.Add(24 * time.Hour)
	return c.eventRepo.GetByPeriod(start, end)
}

// GetForWeek - Получение событий за текущую неделю
func (c *CalendarService) GetForWeek() ([]model.Event, error) {
	//TODO: test and fix
	start := time.Now().Add(-7 * 24 * time.Hour).Round(7 * 24 * time.Hour)
	end := start.Add(7 * 24 * time.Hour)
	return c.eventRepo.GetByPeriod(start, end)
}

// GetForMonth - Получение событий за текущий месяц
func (c *CalendarService) GetForMonth() ([]model.Event, error) {
	//TODO: test and fix
	start := time.Now().Add(-30 * 24 * time.Hour).Round(30 * 24 * time.Hour)
	end := start.Add(30 * 24 * time.Hour)
	return c.eventRepo.GetByPeriod(start, end)
}
