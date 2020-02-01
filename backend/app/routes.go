package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

// AddRoutes configures the routes handlers in the router
func AddRoutes(router *mux.Router, t *TagsStorage) {

	h := &handler{tags: t}

	router.HandleFunc("/", h.sendBasicInfo).Methods("GET")
	router.HandleFunc("/tags", h.getAllExistingTags).Methods("GET")
	router.HandleFunc("/tags/{tag}", h.deleteTagFromAllRepositories).Methods("DELETE")
	router.HandleFunc("/{user}/tags/{tag}", h.getAllRepositoriesWithMatchingTag).Methods("GET")
	router.HandleFunc("/{user}/repos", h.getAllUserStarredRepositoriesAndTags).Methods("GET")
	router.HandleFunc("/{user}/repos/{repo}", h.getRepositoryTags).Methods("GET")
	router.HandleFunc("/{user}/repos/{repo}", h.updateRepositoryTags).Methods("PUT")
	router.HandleFunc("/{user}/repos/{repo}", h.clearRepositoryTags).Methods("DELETE")
	router.HandleFunc("/{user}/repos/{repo}/tags", h.addRepositoryTag).Methods("POST")
	router.HandleFunc("/{user}/repos/{repo}/tags/{tag}", h.removeRepositoryTag).Methods("DELETE")
}

type handler struct {
	tags *TagsStorage
}

type repositoryWithTags struct {
	ID            int64
	GithubID      string
	Name          string
	Description   string
	URL           string
	Language      []string
	Tags          []string
	SuggestedTags []string
}

func (h *handler) sendBasicInfo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *handler) getAllExistingTags(w http.ResponseWriter, r *http.Request) {
	allTags := h.tags.GetAllTags()
	json.NewEncoder(w).Encode(allTags)
}

func (h *handler) getAllRepositoriesWithMatchingTag(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tag := params["tag"]
	user := params["user"]

	maxRepos := GetUserStarredReposCount(user)
	repositories := GetUserStarredRepos(user, maxRepos)
	repos := h.tags.GetReposByTagPattern(tag)
	var filteredRepos []GithubRepository
	for _, starredRepo := range repositories {
		if !contains(repos, starredRepo.ID) {
			continue
		}
		filteredRepos = append(filteredRepos, starredRepo)
	}
	reposWithTags := getReposWithTags(filteredRepos, h.tags)
	json.NewEncoder(w).Encode(reposWithTags)
}

func (h *handler) deleteTagFromAllRepositories(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tag := params["tag"]

	h.tags.DeleteTag(tag)
}

func getReposWithTags(repositories []GithubRepository, tags *TagsStorage) []repositoryWithTags {
	var reposWithTags []repositoryWithTags
	for _, starredRepo := range repositories {
		repoWithTags := getRepositoryWithTags(starredRepo, tags)
		repoWithTags.SuggestedTags = suggestTags(starredRepo, repoWithTags.Tags, tags)
		reposWithTags = append(reposWithTags, repoWithTags)
	}
	return reposWithTags
}

func (h *handler) getAllUserStarredRepositoriesAndTags(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := params["user"]
	maxRepos := GetUserStarredReposCount(user)
	repositories := GetUserStarredRepos(user, maxRepos)
	reposWithTags := getReposWithTags(repositories, h.tags)
	json.NewEncoder(w).Encode(&reposWithTags)
}

func (h *handler) getRepositoryTags(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	repo, err := strconv.ParseInt(params["repo"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	tags := h.tags.GetRepoTags(repo)

	json.NewEncoder(w).Encode(&tags)
}

func (h *handler) updateRepositoryTags(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	repo, err := strconv.ParseInt(params["repo"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var tags []string
	if err := json.NewDecoder(r.Body).Decode(&tags); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h.tags.ClearRepoTags(repo)

	for _, tag := range tags {
		h.tags.AddRepoTag(repo, string(tag))
	}
}

func (h *handler) clearRepositoryTags(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	repo, err := strconv.ParseInt(params["repo"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	h.tags.ClearRepoTags(repo)
}

func (h *handler) addRepositoryTag(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	repo, err := strconv.ParseInt(params["repo"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	tag, _ := ioutil.ReadAll(r.Body)
	h.tags.AddRepoTag(repo, string(tag))

	w.WriteHeader(http.StatusCreated)
}

func (h *handler) removeRepositoryTag(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	repo, err := strconv.ParseInt(params["repo"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	tag, _ := ioutil.ReadAll(r.Body)

	h.tags.RemoveRepoTag(repo, string(tag))
}

func suggestTags(sr GithubRepository, tagsStorage *TagsStorage) []string {
	return tagsStorage.GetAllTags()
}

func convertLanguages(githubRepo GithubRepository) []string {
	var langs []string
	for _, lang := range githubRepo.Languages.Nodes {
		langs = append(langs, lang.Name)
	}
	return langs
}

func getRepositoryWithTags(githubRepo GithubRepository, tags *TagsStorage) repositoryWithTags {
	var repoWithTags repositoryWithTags
	repoWithTags.ID = tags.GetRepoID(githubRepo.ID)
	repoWithTags.GithubID = githubRepo.ID
	repoWithTags.Name = githubRepo.Name
	repoWithTags.Description = githubRepo.Description
	repoWithTags.URL = githubRepo.URL
	repoWithTags.Language = convertLanguages(githubRepo)
	repoWithTags.Tags = tags.GetRepoTags(repoWithTags.ID)
	return repoWithTags
}

func contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
