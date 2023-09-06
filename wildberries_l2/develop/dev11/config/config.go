// Package config - Пакет для работы с конфигурацией
package config

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// Config - структура с параметрами конфигурации сервиса
type Config struct {
	DateFormat string
	ServerPort string
}

// New - загружает конфигурацию из файла
func New(filepath string) (Config, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	cfgMap := readCfgMap(file)

	return Config{
		DateFormat: cfgMap["date_format"],
		ServerPort: cfgMap["port"],
	}, nil
}

func readCfgMap(r io.Reader) map[string]string {
	cfgMap := make(map[string]string)

	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if strings.HasPrefix(line, "#") {
			continue
		}

		pair := strings.SplitN(line, "=", 2)
		if len(pair) == 2 {
			key := strings.TrimSpace(pair[0])
			value := strings.TrimSpace(pair[1])
			cfgMap[key] = value
		}
	}

	return cfgMap
}
