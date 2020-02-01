package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// AddMiddlewares configures the middlewares in the router
func AddMiddlewares(router *mux.Router) {
	router.Use(logURLMiddleware)
}

func logURLMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println(r.Method, r.URL)
        h.ServeHTTP(w, r)
    })
}

func headerMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("accept", "application/json")
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
        h.ServeHTTP(w, r)
    })
}
