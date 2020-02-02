# GithubStars

The projects in this repository are for learning purposes. They are designed to allow the user to browse GitHub repositories that have been starred by the provided username and mark tags on them. They also allow the user to filter the browse results by tag name.

- [backend](./backend) project is an REST API server written in Go that consumes the GitHub v4 GraphQL API and stores the data in a local SQLite database.
- [frontend](./frontend) poject is a Vue app that consumes the GitHub v3 REST API and stores the data in the user's local storage

Both projects contain a docker image to ease the application execution.
