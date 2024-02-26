package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", "localhost:8000", "The address of the sever")
	flag.Parse()

	fmt.Fprintln(os.Stderr, "Server starting on: ", *addr)

	shortr := http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "testing server out...")
		},
	)

	if err := http.ListenAndServe(*addr, shortr); !errors.Is(err, http.ErrServerClosed) {
		fmt.Fprintln(os.Stderr, "err, server closed unexpectedly: %w", err)
	}
}
