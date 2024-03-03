package short_server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/marktlinn/gShort/bite"
	shortsvc "github.com/marktlinn/gShort/short_svc"
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

// Parses the URL & short key from the Request
// and stores them in a new Link.
func handleShorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(
			w,
			"method not allowed: should be POST",
			http.StatusMethodNotAllowed,
		)
		return
	}

	ln := shortsvc.Link{
		Key: r.FormValue("key"),
		URL: r.FormValue("url"),
	}
	if err := shortsvc.Create(r.Context(), ln); err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "go")
}

// The route path is stripped from the Request in the process.
func handleResolve(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[len(resolvedRoute):]

	ln, err := shortsvc.Retrieve(r.Context(), key)
	if err != nil {
		handleError(w, err)
		return
	}
	http.Redirect(w, r, ln.URL, http.StatusFound)
}

func handleHealthRoute(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "OK")
}

func handleError(w http.ResponseWriter, err error) {
	switch {
	case err == nil:
		return
	case errors.Is(err, bite.ErrInvalidRequest):
		http.Error(w, err.Error(), http.StatusBadRequest)
	case errors.Is(err, bite.ErrNotExists):
		http.Error(w, err.Error(), http.StatusNotFound)
	case errors.Is(err, bite.ErrExists):
		http.Error(w, err.Error(), http.StatusConflict)
	}
}
