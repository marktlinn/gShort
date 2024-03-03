package httpio

import (
	"context"
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

// Extracts the Value from the Server's Context and uses type assertion
// to log with the Server's Logger, if the Server and Logger exist
func Log(ctx context.Context, format string, args ...any) {
	s, _ := ctx.Value(http.ServerContextKey).(*http.Server)
	if s == nil || s.ErrorLog == nil {
		return
	}
	s.ErrorLog.Printf(format, args...)
}
