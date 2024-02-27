package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	var (
		addr    = flag.String("addr", "localhost:8000", "The address of the sever")
		timeout = flag.Duration(
			"t",
			10*time.Second,
			"sets the duration for which the server will wait for requests before a timeout",
		)
	)
	flag.Parse()

	fmt.Fprintln(os.Stderr, "Server starting on: ", *addr)

	shortr := http.HandlerFunc(
		func(w http.ResponseWriter, _ *http.Request) {
			fmt.Fprintln(w, "testing server out...")
		},
	)

	server := &http.Server{
		Addr:        *addr,
		Handler:     http.TimeoutHandler(shortr, *timeout, "timeout"),
		ReadTimeout: *timeout,
	}

	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		fmt.Fprintln(os.Stderr, "err, server closed unexpectedly: %w", err)
	}
}
