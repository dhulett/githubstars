FORMAT: 1A

# GithubStars

GithubStars is a simple API allowing consumers to browse Github repositories starred by the provided user, add tags to the repositories and browse searching by tags.
You can view this documentation over at [Apiary](https://githubstars.docs.apiary.io/).

## GithubStars API Root [/]

### Test server connection [GET]

+ Response 204 (application/json)

## Group Tags

Resources related to Tags in the API

### Tag [/tags/{tag}]

+ Parameter
    + tag: "tag" (required, string) - The tag created by the user

#### Delete Tag [DELETE]

+ Response 204 (application/json)

### Tags Collection [/tags]

#### List all tags [GET]

+ Response 200 (application/json)

        [
            "tag1",
            "tag2"
        ]

### Tags in user repositories [/tags/{tag}/{user}]

+ Parameters
    + tag: "tag" - Tag string
    + user: "username" - Username of the Github user

#### Search repositories by tag [GET]

+ Response 200 (application/json)

        [
            {
                "ID": 1,
                "GithubID": "github_id",
                "Name": "RepositoryName",
                "Description": "Repository description",
                "Owner": "owner",
                "URL": "https://github.com/owner/RepositoryName",
                "Language": [
                    "Language1",
                    "Language2"
                ],
                "Tags": [
                    "tag1",
                    "tag2"
                ],
                "SuggestedTags": [
                    "suggested1",
                    "suggested2"
                ]
            }
        ]

## Group Repo

Resources related to Repos in the API

### Repo [/{user}/repos/{repo}]

A Repo object represents a Github repository and has the following attributes:

+ Attributes:
    + ID - Internal ID to use in the API calls
    + GithubID
    + Name
    + Description
    + Owner
    + URL
    + Language
    + Tags
    + SuggestedTags

+ Parameters
    + user: username (required, string) - Username of the Github user
    + repo: 1 (required, string) - Internal ID of the repository

#### Get Repo details [GET]

+ Response 200 (application/json)

        {
            "ID": 1,
            "GithubID": "github_id",
            "Name": "RepositoryName",
            "Description": "Repository description",
            "Owner": "owner",
            "URL": "https://github.com/owner/RepositoryName",
            "Language": [
                "Language1",
                "Language2"
            ],
            "Tags": [
                "tag1",
                "tag2"
            ],
            "SuggestedTags": [
                "suggested1",
                "suggested2"
            ]
        }

+ Response 204 (application/json)

+ Response 400 (application/json)

### Repo Collection [/{user}/repos]

+ Parameters
    + user: username (required, string) - Username of the Github user

#### List all user starred Repos [GET]

+ Response 200 (application/json)

        [
            {
                "ID": 1,
                "GithubID": "github_id_1",
                "Name": "Repo1Name",
                "Description": "Repo1 description",
                "Owner": "repo1owner",
                "URL": "https://github.com/repo1owner/Repo1Name",
                "Language": [
                    "Lang1"
                ],
                "Tags": [
                    "tag1"
                ],
                "SuggestedTags": [
                    "suggested"
                ]
            }
        ]

### Repo Tags Collection [/{user}/repos/{repo}/tags]

+ Parameters
    + user: username (required, string) - Username of the Github user
    + repo: 1 (required, string) - Internal ID of the repository

#### Create new tag [POST]

+ Request

        "tag"

+ Response 201 (application/json)

+ Response 400 (application/json)

#### Get Repo Tags [GET]

+ Response 200 (application/json)

        [
            "tag1"
        ]

+ Response 400 (application/json)

#### Update Repo Tags [PATCH]

+ Request

        [
            "tag1",
            "tag2"
        ]

+ Response 204 (application/json)

+ Response 400 (application/json)

#### Clear Repo Tags [DELETE]

+ Response 204 (application/json)

+ Response 400 (application/json)

### Repo Tag [/{user}/reposq{repo}/tags/{tag}]

+ Parameters
    + user: "username" (required, string) - Username of the Github user
    + repo: 1 (required, string) - Internal ID of the repository
    + tag: "tag" (required, string) - The tag to mark repositories

#### Remove Tag from Repo [DELETE]

+ Response 204 (application/json)

+ Response 400 (application/json)
