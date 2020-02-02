package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

// AddRoutes configures the routes handlers in the router
func AddRoutes(router *mux.Router, t *TagsStorage, g *GithubGraphQLClient) {

	h := &handler{tags: t, github: g}

	router.HandleFunc("/", h.sendBasicInfo).Methods("GET")
	router.HandleFunc("/tags", h.getAllExistingTags).Methods("GET")
	router.HandleFunc("/tags/{tag}", h.deleteTagFromAllRepositories).Methods("DELETE")
	router.HandleFunc("/tags/{tag}/{user}", h.getAllRepositoriesWithMatchingTag).Methods("GET")
	router.HandleFunc("/{user}/repos", h.getAllUserStarredRepositoriesAndTags).Methods("GET")
	router.HandleFunc("/{user}/repos/{repo}", h.getStarredRepoDetails).Methods("GET")
	router.HandleFunc("/{user}/repos/{repo}/tags", h.addRepositoryTag).Methods("POST")
	router.HandleFunc("/{user}/repos/{repo}/tags", h.getRepositoryTags).Methods("GET")
	router.HandleFunc("/{user}/repos/{repo}/tags", h.updateRepositoryTags).Methods("PUT")
	router.HandleFunc("/{user}/repos/{repo}/tags", h.clearRepositoryTags).Methods("DELETE")
	router.HandleFunc("/{user}/repos/{repo}/tags/{tag}", h.removeRepositoryTag).Methods("DELETE")
}

type handler struct {
	tags *TagsStorage
	github *GithubGraphQLClient
}

type repositoryWithTags struct {
	ID            int64
	GithubID      string
	Name          string
	Description   string
	Owner         string
	URL           string
	Languages     []string
	Tags          []string
	SuggestedTags []string
}

func (h *handler) sendBasicInfo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) getAllExistingTags(w http.ResponseWriter, r *http.Request) {
	allTags := h.tags.GetAllTags()
	if len(allTags) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	json.NewEncoder(w).Encode(&allTags)
}

func (h *handler) getAllRepositoriesWithMatchingTag(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tag := params["tag"]
	user := params["user"]

	maxRepos := h.github.GetUserStarredReposCount(user)
	repositories := h.github.GetUserStarredRepos(user, maxRepos)
	repos := h.tags.GetReposByTagPattern(tag)
	var filteredRepos []GithubRepository
	for _, starredRepo := range repositories {
		if !contains(repos, starredRepo.ID) {
			continue
		}
		filteredRepos = append(filteredRepos, starredRepo)
	}
	reposWithTags := getReposWithTags(filteredRepos, h.tags)
	if reposWithTags == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	json.NewEncoder(w).Encode(&reposWithTags)
}

func (h *handler) deleteTagFromAllRepositories(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tag := params["tag"]

	h.tags.DeleteTag(tag)

	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) getAllUserStarredRepositoriesAndTags(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := params["user"]
	maxRepos := h.github.GetUserStarredReposCount(user)
	repositories := h.github.GetUserStarredRepos(user, maxRepos)
	reposWithTags := getReposWithTags(repositories, h.tags)
	if reposWithTags == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	json.NewEncoder(w).Encode(&reposWithTags)
}

func (h *handler) getRepositoryTags(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	repo, err := strconv.ParseInt(params["repo"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tags := h.tags.GetRepoTags(repo)
	if tags == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	json.NewEncoder(w).Encode(&tags)
}

func (h *handler) getStarredRepoDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := params["user"]
	repo, err := strconv.ParseInt(params["repo"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	githubID := h.tags.GetRepoGithubID(repo)

	if githubID != "" {
		maxRepos := h.github.GetUserStarredReposCount(user)
		repositories := h.github.GetUserStarredRepos(user, maxRepos)
		for _, githubRepo := range repositories {
			if githubRepo.ID == githubID {
				repoWithTags := getRepositoryWithTags(githubRepo, h.tags)
				json.NewEncoder(w).Encode(repoWithTags)
				return
			}
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) updateRepositoryTags(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	repo, err := strconv.ParseInt(params["repo"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h.tags.ClearRepoTags(repo)
}

func (h *handler) addRepositoryTag(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	repo, err := strconv.ParseInt(params["repo"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tag := params["tag"]
	h.tags.RemoveRepoTag(repo, string(tag))
	w.WriteHeader(http.StatusNoContent)
}

func suggestTags(githubRepo GithubRepository, repoTags []string, tagsStorage *TagsStorage) []string {
	var suggestedTags []string
	languages := convertLanguages(githubRepo)
	for _, lang := range languages {
		if !contains(suggestedTags, lang) && !contains(repoTags, lang) {
			suggestedTags = append(suggestedTags, lang)
		}
	}

	for _, tag := range tagsStorage.GetAllTags() {
		if !contains(suggestedTags, tag) && !contains(repoTags, tag) {
			suggestedTags = append(suggestedTags, tag)
		}
	}

	if !contains(suggestedTags, githubRepo.Owner.Login) && !contains(repoTags, githubRepo.Owner.Login) {
		suggestedTags = append(suggestedTags, githubRepo.Owner.Login)
	}

	return suggestedTags
}

func convertLanguages(githubRepo GithubRepository) []string {
	var langs []string
	for _, lang := range githubRepo.Languages.Nodes {
		langs = append(langs, lang.Name)
	}
	return langs
}

func getReposWithTags(repositories []GithubRepository, tags *TagsStorage) []repositoryWithTags {
	var reposWithTags []repositoryWithTags
	for _, githubRepo := range repositories {
		repoWithTags := getRepositoryWithTags(githubRepo, tags)
		reposWithTags = append(reposWithTags, repoWithTags)
	}
	return reposWithTags
}

func getRepositoryWithTags(githubRepo GithubRepository, tags *TagsStorage) repositoryWithTags {
	var repoWithTags repositoryWithTags
	repoWithTags.ID = tags.GetRepoID(githubRepo.ID)
	repoWithTags.GithubID = githubRepo.ID
	repoWithTags.Name = githubRepo.Name
	repoWithTags.Owner = githubRepo.Owner.Login
	repoWithTags.Description = githubRepo.Description
	repoWithTags.URL = githubRepo.URL
	repoWithTags.Languages = convertLanguages(githubRepo)
	repoWithTags.Tags = tags.GetRepoTags(repoWithTags.ID)
	repoWithTags.SuggestedTags = suggestTags(githubRepo, repoWithTags.Tags, tags)
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
