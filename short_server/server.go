package short_server

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	shorteningRoute  = "/shorten"
	resolvedRoute    = "/r/"
	healthCheckRoute = "/health"
)

type Server struct {
	// TODO
}

func (s *Server) RegisterRoutes() {
	// TODO
}

func (s *Server) SeverHTTP(w http.ResponseWriter, r *http.Request) {
	switch p := r.URL.Path; {
	case p == shorteningRoute:
		handleShorten(w, r)
	case strings.HasPrefix(p, resolvedRoute):
		handleResolve(w, r)
	case p == healthCheckRoute:
		handleHealthRoute(w, r)
	default:
		http.NotFound(w, r)
	}
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
