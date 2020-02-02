package main

import (
	"context"
	"github.com/machinebox/graphql"
	"log"
)

const githubGraphQL = "https://api.github.com/graphql"

// GithubGraphQLClient handles requests to the Github v4 GraphQL API
type GithubGraphQLClient struct {
	client *graphql.Client
	auth string
}

// GetGithubClient returns a Github API client
func GetGithubClient() *GithubGraphQLClient {
	client := &GithubGraphQLClient{
			client: graphql.NewClient(githubGraphQL),
			auth: "",
		}
	return client
}

// GithubRepository holds information about a starred repository
type GithubRepository struct {
	ID          string
	Name        string
	Description string
	URL         string
	Languages   struct {
		Nodes   []ProgrammingLanguage
	}
	Owner       struct {
		Login    string
	}
}

// ProgrammingLanguage holds information about the programming languange marked on the repository
type ProgrammingLanguage struct {
	Name string
}

type apiResponse struct {
	User struct {
		StarredRepositories struct {
			Nodes []GithubRepository
		}
	}
}

// GetUserStarredRepos returns all the starred repos for the user
func (github *GithubGraphQLClient) GetUserStarredRepos(user string, maxRepos int) []GithubRepository {
	req := getStarredReposRequest(user, maxRepos, 1)
	req.Header.Add("Authorization", github.getAuthorizationHeader())
	ctx := context.Background()
	var res apiResponse
	if err := github.client.Run(ctx, req, &res); err != nil {
		log.Fatal(err)
	}
	return res.User.StarredRepositories.Nodes
}

type apiCountResponse struct {
	User struct {
		StarredRepositories struct {
			TotalCount int
		}
	}
}

// GetUserStarredReposCount returns the number of repos starred by the user
func (github *GithubGraphQLClient) GetUserStarredReposCount(user string) int {
	req := getStarredReposCountRequest(user)
	req.Header.Add("Authorization", github.getAuthorizationHeader())
	ctx := context.Background()
	var res apiCountResponse
	if err := github.client.Run(ctx, req, &res); err != nil {
		log.Fatal(err)
	}
	return res.User.StarredRepositories.TotalCount
}

func (github *GithubGraphQLClient) getAuthorizationHeader() string {
	return "Bearer " + github.auth
}

func getStarredReposCountRequest(user string) *graphql.Request {
	req := graphql.NewRequest(`query ($user: String!) {
		user(login:$user){
			starredRepositories{
				totalCount
			}
		}
	}`)

	req.Var("user", user)

	return req
}

func getStarredReposRequest(user string, maxRepos int, maxLanguages int) *graphql.Request {
	req := graphql.NewRequest(`query ($user: String!, $maxRepos: Int!, $maxLanguages: Int!) {
		user(login:$user) {
			starredRepositories(first:$maxRepos) {
				totalCount
				nodes {
					id
					name
					description
					url
					languages(first:$maxLanguages) {
						nodes {
							name
						}
					}
					owner {
           				login
					}
				}
			}
		}
	}`)

	req.Var("user", user)
	req.Var("maxRepos", maxRepos)
	req.Var("maxLanguages", maxLanguages)

	return req
}
