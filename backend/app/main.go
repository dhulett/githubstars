package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const port = ":8000"

func main() {
	router := mux.NewRouter()
	tags := GetTagsStorage("./githubstars.db")
	AddRoutes(router, tags)
	AddMiddlewares(router)
	fmt.Println("Listening on", port, "...")
	log.Fatal(http.ListenAndServe(port, router))
}
