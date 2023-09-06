// Package server - пакет с реализацией сервера для работы с календарем
package server

import (
	"errors"
	"fmt"
	"github.com/Generat17/dev11/model"
	"github.com/Generat17/dev11/service"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

// HTTPServer - Сервер для работы с календарем
type HTTPServer struct {
	mux        *http.ServeMux
	calendar   *service.CalendarService
	serializer model.EventSerializer
}

// NewHTTPServer - создает готовый к работе сервер
func NewHTTPServer(
	calendar *service.CalendarService,
	serializer model.EventSerializer,
) *HTTPServer {
	mux := http.NewServeMux()
	s := HTTPServer{
		mux:        mux,
		calendar:   calendar,
		serializer: serializer,
	}

	mux.Handle("/create_event",
		loggerMiddleware(
			strictMethodMiddleware(
				map[string]http.Handler{
					"POST": http.HandlerFunc(s.createEventHandler),
				},
			),
		),
	)

	mux.Handle("/update_event",
		loggerMiddleware(
			strictMethodMiddleware(
				map[string]http.Handler{
					"POST": http.HandlerFunc(s.updateEventHandler),
				},
			),
		),
	)

	mux.Handle("/delete_event",
		loggerMiddleware(
			strictMethodMiddleware(
				map[string]http.Handler{
					"POST": http.HandlerFunc(s.deleteEventHandler),
				},
			),
		),
	)

	mux.Handle("/events_for_day",
		loggerMiddleware(
			strictMethodMiddleware(
				map[string]http.Handler{
					"GET": http.HandlerFunc(s.getForDayHandler),
				},
			),
		),
	)

	mux.Handle("/events_for_week",
		loggerMiddleware(
			strictMethodMiddleware(
				map[string]http.Handler{
					"GET": http.HandlerFunc(s.getForWeekHandler),
				},
			),
		),
	)

	mux.Handle("/events_for_month",
		loggerMiddleware(
			strictMethodMiddleware(
				map[string]http.Handler{
					"GET": http.HandlerFunc(s.getForMonthHandler),
				},
			),
		),
	)

	return &s
}

// ServeHTTP - запускает сервер на заданном адресе
func (s *HTTPServer) ServeHTTP(addr string) error {
	return http.ListenAndServe(addr, s.mux)
}

func (s *HTTPServer) createEventHandler(w http.ResponseWriter, r *http.Request) {
	var dto service.EventCreateDTO
	err := r.ParseForm()
	if err == nil {
		dto, err = parseCreateEventDTO(r.Form)
	}
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	event, err := s.calendar.CreateEvent(dto)

	if err != nil {
		log.Println(err)
		writeErr(w, err)
		return
	}

	bytes, err := s.serializer.Serialize(event)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	writeResult(w, bytes)
}

func (s *HTTPServer) updateEventHandler(w http.ResponseWriter, r *http.Request) {
	var dto service.EventUpdateDTO
	err := r.ParseForm()
	if err == nil {
		dto, err = parseUpdateEventDTO(r.Form)
	}
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	event, err := s.calendar.UpdateEvent(dto)

	if err != nil {
		log.Println(err)
		writeErr(w, err)
		return
	}

	bytes, err := s.serializer.Serialize(event)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	writeResult(w, bytes)
}

func (s *HTTPServer) deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	var id int
	err := r.ParseForm()
	if err == nil {
		id, err = strconv.Atoi(r.Form.Get("id"))
	}
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.calendar.DeleteEvent(model.EventID(id))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *HTTPServer) getForDayHandler(w http.ResponseWriter, r *http.Request) {
	events, err := s.calendar.GetForDay()
	if err != nil {
		log.Println(err)
		writeErr(w, err)
		return
	}

	bytes, err := s.serializer.SerializeArray(events)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	writeResult(w, bytes)
}

func (s *HTTPServer) getForWeekHandler(w http.ResponseWriter, r *http.Request) {
	events, err := s.calendar.GetForWeek()
	if err != nil {
		log.Println(err)
		writeErr(w, err)
		return
	}

	bytes, err := s.serializer.SerializeArray(events)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	writeResult(w, bytes)
}

func (s *HTTPServer) getForMonthHandler(w http.ResponseWriter, r *http.Request) {
	events, err := s.calendar.GetForMonth()
	if err != nil {
		log.Println(err)
		writeErr(w, err)
		return
	}

	bytes, err := s.serializer.SerializeArray(events)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	writeResult(w, bytes)
}

func parseCreateEventDTO(form url.Values) (service.EventCreateDTO, error) {
	if !form.Has("date") || !form.Has("description") {
		return service.EventCreateDTO{}, errors.New("bad request")
	}

	return service.EventCreateDTO{
		Date:        form.Get("date"),
		Description: form.Get("description"),
	}, nil
}

func parseUpdateEventDTO(form url.Values) (service.EventUpdateDTO, error) {
	if !form.Has("id") || !form.Has("date") && !form.Has("description") {
		return service.EventUpdateDTO{}, errors.New("bad request")
	}

	var dto service.EventUpdateDTO
	id, err := strconv.Atoi(form.Get("id"))
	if err != nil {
		return service.EventUpdateDTO{}, errors.New("bad id")
	}
	dto.ID = model.EventID(id)

	if form.Has("date") {
		date := form.Get("date")
		dto.Date = &date
	}
	if form.Has("description") {
		description := form.Get("description")
		dto.Description = &description
	}
	return dto, nil
}

func writeErr(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusServiceUnavailable)
	fmt.Fprintf(w, `{"error":"%s"}`, err.Error())
}

func writeResult(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"result":%s}`, string(bytes))
}
