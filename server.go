// Package singlereq provides HTTP servers that listens for a few requests.
package fewrequests

import (
	"context"
	"errors"
	"log"
	"net/http"
)

// ListenAndServeOnce starts an http.Server that shutdowns after handling one
// request.
//
// The handler is typically nil, in which case the http.DefaultServerMux is used.
func ListenAndServeOnce(addr string, h http.Handler) error {
	return ListenAndServeN(1, addr, h)
}

// ListenAndServe starts an http.Server that shutdowns after handling N
// requests.
//
// The handler is typically nil, in which case the http.DefaultServerMux is used.
func ListenAndServeN(N int, addr string, h http.Handler) error {
	if N == 0 {
		return errors.New("cannot serve for 0 request")
	}

	counter := 0
	s := http.Server{
		Addr: addr,
	}

	if h == nil {
		h = http.DefaultServeMux
	}

	// Wrapping the handler in a new one and register it.
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		counter++
		if counter >= N {
			// Shutdown from another go routine.
			go func() {
				if err := s.Shutdown(context.Background()); err != nil {
					log.Fatalf("server shutdown: %v", err)
				}
			}()
			return
		}
	})

	s.Handler = handler

	return s.ListenAndServe()
}
