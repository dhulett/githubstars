package main

import (
	"context"
	"github.com/machinebox/graphql"
	"log"
)

// StarredRepository holds information about a starred repository
type StarredRepository struct {
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
		TotalCount          int64
		StarredRepositories struct {
			Nodes []StarredRepository
		}
	}
}

// GetUserStarredRepos returns all the starred repos for the user
func GetUserStarredRepos(user string) []StarredRepository {
	client := graphql.NewClient("https://api.github.com/graphql")
	req := getStarredReposRequest(user)
	req.Header.Add("Authorization", "Bearer 03b8cd62c73d8acad36e6a5f7ba5bc8c907c2eb7")
	ctx := context.Background()
	var res apiResponse
	if err := client.Run(ctx, req, &res); err != nil {
		log.Fatal(err)
	}
	return res.User.StarredRepositories.Nodes
}

func getStarredReposRequest(user string) *graphql.Request {
	req := graphql.NewRequest(`query ($user: String!) {
		user(login:$user){
			starredRepositories(first:10){
				totalCount
				nodes
				{
					id
					name
					description
					url
					languages(first:10)
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

	return req
}
