package short_server

import (
	"fmt"
	"net/http"
)

const (
	shorteningRoute  = "/shorten"
	resolvedRoute    = "/r/"
	healthCheckRoute = "/health"
)

type Server struct {
	http.Handler
}

func (s *Server) RegisterRoutes() {
	mux := http.NewServeMux()
	mux.HandleFunc(shorteningRoute, handleShorten)
	mux.HandleFunc(resolvedRoute, handleResolve)
	mux.HandleFunc(healthCheckRoute, handleHealthRoute)
	s.Handler = mux
}

func handleShorten(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "go")
}

func handleResolve(w http.ResponseWriter, r *http.Request) {
	const uri = "https://go.dev"
	http.Redirect(w, r, uri, http.StatusFound)
}

func handleHealthRoute(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "OK")
}
