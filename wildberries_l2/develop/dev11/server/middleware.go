// Package server - пакет с реализацией сервера для работы с календарем
package server

import (
	"log"
	"net/http"
)

func strictMethodMiddleware(nextMap map[string]http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		method := r.Method

		if next, ok := nextMap[method]; ok {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func loggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path, r.Header)
		next.ServeHTTP(w, r)
	}
}

// func enforceJSONMiddleware(next http.Handler) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		contentType := r.Header.Get("Content-Type")

// 		if contentType == "" {
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}

// 		mediaType, _, err := mime.ParseMediaType(contentType)
// 		if err != nil || mediaType != "application/json" {
// 			w.WriteHeader(http.StatusUnsupportedMediaType)
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	}
// }
