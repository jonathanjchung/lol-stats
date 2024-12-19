package middleware

import (
	"log"
	"net/http"
	"time"
)

type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *ResponseWriter) WriteHeader(status int) {
	rw.statusCode = status
	rw.ResponseWriter.WriteHeader(status)
}

func LogMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()
			wrapper := &ResponseWriter{ResponseWriter: w}
			next.ServeHTTP(wrapper, r)
			elapsedTime := time.Since(startTime)
			log.Printf("[%s] [%s] [%s]\n", r.Method, r.URL.Path, elapsedTime)
			log.Printf("Status Code: [%d]", wrapper.statusCode)
		},
	)
}
