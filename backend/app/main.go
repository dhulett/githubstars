package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const port = ":8000"

func main() {
	tags := GetTagsStorage("./githubstars.db")
	githubClient := GetGithubClient()
	router := mux.NewRouter()
	AddRoutes(router, tags, githubClient)
	AddMiddlewares(router, githubClient)
	fmt.Println("Listening on", port, "...")
	log.Fatal(http.ListenAndServe(port, router))
}
