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
	githubClient := GetGithubClient("7d78b293e54d870a1b320190b0afff72e4db65f8")
	router := mux.NewRouter()
	AddRoutes(router, tags, githubClient)
	AddMiddlewares(router)
	fmt.Println("Listening on", port, "...")
	log.Fatal(http.ListenAndServe(port, router))
}
