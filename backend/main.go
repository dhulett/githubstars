package main

import (
	"fmt"
	"context"
	"log"
	"github.com/machinebox/graphql"
	// "github.com/gorilla/mux"
	// "encoding/json"
	//	"net/http"
	//	"github.com/gorilla/mux"
	//	"github.com/graphql-go/graphql"
)

func main() {

	respData := getUserStarredRepos("dhulett")
	fmt.Printf("Results: %v\n", respData.User.StarredRepositories.Nodes)
}

func getUserStarredRepos(user string) Response {
	client := graphql.NewClient("https://api.github.com/graphql")
	req := getStarredReposRequest(user)
	req.Header.Add("Authorization", "Bearer 03b8cd62c73d8acad36e6a5f7ba5bc8c907c2eb7")
	ctx := context.Background()
	var respData Response
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
	}
	return respData
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

type Response struct {
	User struct {
		StarredRepositories struct {
			TotalCount int64
			Nodes      []struct {
				Id          string
				Name        string
				Description string
				Url         string
				Languages   struct {
					Nodes []struct {
						Name string
					}
				}
			}
		}
	}
}
