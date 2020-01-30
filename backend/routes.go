package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	// "context"
	// "log"
	// "encoding/json"
	// "github.com/machinebox/graphql"
	"github.com/gorilla/mux"
)

// PrepareRouter configures the routes handlers in the router
func PrepareRouter(router *mux.Router) {

	router.HandleFunc("/{user}", getAllUserStarredRepositoriesAndTags).Methods("GET")
	router.HandleFunc("/{user}/tags", getAllExistingTags).Methods("GET")
	router.HandleFunc("/{user}/tags/{tags}", getAllRepositoriesWithTags).Methods("GET")
	router.HandleFunc("/{user}/repos/{repo}", getRepositoryTags).Methods("GET")
	router.HandleFunc("/{user}/repos/{repo}/{tags}", addRepositoryTag).Methods("PUT")
	router.HandleFunc("/{user}/repos/{repo}/{tags}", removeRepositoryTag).Methods("DELETE")

}

func getAllUserStarredRepositoriesAndTags(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetAllUserStarredRepositoriesAndTags")

	params := mux.Vars(r)
	user := params["user"]
	maxRepos := GetUserStarredReposCount(user)
	repositories := GetUserStarredRepos(params["user"], maxRepos)

	json.NewEncoder(w).Encode(&repositories)
}

func getAllExistingTags(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetAllExistingTags")
}

func getAllRepositoriesWithTags(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetAllRepositoriesWithTags")
}

func getRepositoryTags(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetRepositoryTags")
}

func addRepositoryTag(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AddRepositoryTag")
}

func removeRepositoryTag(w http.ResponseWriter, r *http.Request) {
	fmt.Println("RemoveRepositoryTag")
}
