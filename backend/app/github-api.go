package main

import (
	"context"
	"github.com/machinebox/graphql"
	"log"
)

// GithubRepository holds information about a starred repository
type GithubRepository struct {
	ID          string
	Name        string
	Description string
	URL         string
	Languages   struct {
		Nodes []ProgrammingLanguage
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
func GetUserStarredRepos(user string, maxRepos int) []GithubRepository {
	client := graphql.NewClient("https://api.github.com/graphql")

	req := getStarredReposRequest(user, maxRepos, 1)
	req.Header.Add("Authorization", "Bearer 03b8cd62c73d8acad36e6a5f7ba5bc8c907c2eb7")
	ctx := context.Background()
	var res apiResponse
	if err := client.Run(ctx, req, &res); err != nil {
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
func GetUserStarredReposCount(user string) int {
	client := graphql.NewClient("https://api.github.com/graphql")

	req := getStarredReposCountRequest(user)
	req.Header.Add("Authorization", "Bearer 03b8cd62c73d8acad36e6a5f7ba5bc8c907c2eb7")
	ctx := context.Background()
	var res apiCountResponse
	if err := client.Run(ctx, req, &res); err != nil {
		log.Fatal(err)
	}
	return res.User.StarredRepositories.TotalCount
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
		user(login:$user){
			starredRepositories(first:$maxRepos){
				totalCount
				nodes
				{
					id
					name
					description
					url
					languages(first:$maxLanguages)
					{
						nodes
						{
							name
						}
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
