# GithubStars Server

A server that exposes a REST API to allow the user to fetch information about repositories that a provided github user has starred and mark tags to the repositories.

- The API is described in [GithubStarsAPI](GithubStarsAPI.md)
- The data is stored in a local SQLite database.
- The user starred repositories information is fetched from the GitHub server using the GitHub v4 API (GraphQL)

## Getting Started

### Run with Go:

Requires:
+ [Go](https://golang.org/) version 1.13 or above (Go Modules)
+ GCC Compiler referenced in the environent %PTH% (to build sqlite3 driver dependency)

Execute following commands to run the application
```shell
$ cd ./app
$ go run .
```

### or

### Run with docker:

Requires:
+ [Docker](https://www.docker.com/)

Execute following commands to run the application
```shell
$ docker run -v ./app/:/go/githubstars -p 8000:8000 --rm -it $(docker build -q .)
```

## Built With

+ [GoLang](https://golang.org/)

## Authors

* **Deon Hulett** - [dhulett](https://github.com/dhulett)
