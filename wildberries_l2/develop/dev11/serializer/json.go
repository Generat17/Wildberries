// Package serializer - пакет для сериализации событий
package serializer

import (
	"encoding/json"
	"github.com/Generat17/dev11/model"
	"strconv"
	"strings"
)

// JSONEventSerializer - сериализатор событий в JSON формате
type JSONEventSerializer struct {
	dateFormat string
}

var _ model.EventSerializer = &JSONEventSerializer{}

// NewJSONEventSerializer - конструктор
func NewJSONEventSerializer(dateFormat string) *JSONEventSerializer {
	return &JSONEventSerializer{
		dateFormat: dateFormat,
	}
}

// Serialize implements model.EventSerializer.
func (s *JSONEventSerializer) Serialize(event model.Event) ([]byte, error) {
	sb := strings.Builder{}

	sb.WriteString(`{"id":`)
	sb.WriteString(strconv.Itoa(int(event.ID)))
	sb.WriteString(`,"date":"`)
	sb.WriteString(event.Date.Format(s.dateFormat))
	sb.WriteString(`","description":`)

	desc, err := json.Marshal(event.Description)
	if err != nil {
		return nil, err
	}
	sb.WriteString(string(desc))
	sb.WriteString(`}`)

	bytes := []byte(sb.String())
	return bytes, nil
}

// SerializeArray implements model.EventSerializer.
func (s *JSONEventSerializer) SerializeArray(events []model.Event) ([]byte, error) {
	bytes := make([][]byte, len(events))

	for i, event := range events {
		b, err := s.Serialize(event)
		if err != nil {
			return nil, err
		}
		bytes[i] = b
	}

	sb := strings.Builder{}
	sb.WriteString("[")
	for i, b := range bytes {
		sb.WriteString(string(b))
		if i < len(bytes)-1 {
			sb.WriteString(",")
		}
	}
	sb.WriteString("]")

	return []byte(sb.String()), nil
}
