package httpio

import (
	"log"
	"net/http"
	"time"
)

// LoggerMiddleware logs the Method, URL Path RemoteAddr
// and time taken to process a received HTTP Request.
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		end := time.Since(start)
		log.Printf("%s %s %s %v", r.Method, r.URL.Path, r.RemoteAddr, end)
	})
}
