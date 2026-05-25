package http_server

import (
	"fmt"
	"net/http"
	"time"
)

func exampleMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		f.ServeHTTP(w, r)
		end := time.Since(start)
		fmt.Println(end)
	})
}
