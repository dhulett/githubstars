FORMAT: 1A

# GithubStars


# GithubStars API Root [/]

## Retrieve the Entry Point [GET]

+ Response 200 (application/json)

        {
            "url": "/repo"
        }

## Group Repos

## Repo [/repos/{id}]

A Repo object holds information about a Github repository.

+ Parameters
    + id

+ Attributes
    + id (string, required)
    + name (string, required)
    + description (string, required)
    + repoUrl (string, required) - The URL of the Github repository
    + language (string, required)
    + tags (array[string], required) - An array of Tag objects.
    + url (string, required) - The current URL in the API

### Retrieve a Repo [GET]
Retrieves the details of the repository with the given id.

+ Response 200 (application/json)
    + Attributes (Repo)

## Repos [/repos{?limit}]

+ Attributes (array[Repo])

### List all Repos [GET]
Retrieves all starred repositories for the current user.

+ Response 200 (application/json)
    + Attributes (Repos)
