package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// AddMiddlewares configures the middlewares in the router
func AddMiddlewares(router *mux.Router, ghClient *GithubGraphQLClient) {
	router.Use(logURLMiddleware)
	router.Use(headerMiddleware)

	ghAuth := &githubAuth{githubClient: ghClient}
	router.Use(ghAuth.githubAuthMiddleware)
}

func logURLMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method, r.URL)
		h.ServeHTTP(w, r)
	})
}

func headerMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		h.ServeHTTP(w, r)
	})
}

type githubAuth struct {
	githubClient *GithubGraphQLClient
}

func (ga *githubAuth) githubAuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ga.githubClient.auth = r.Header.Get("authorization")
		h.ServeHTTP(w, r)
	})
}
