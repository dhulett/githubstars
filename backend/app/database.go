package main

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// TagsStorage holds the database object
type TagsStorage struct {
	db *sql.DB
}

// GetAllTags retrieves all tags from the database
func (t *TagsStorage) GetAllTags() []string {
	tags, err := t.db.Query("SELECT tag FROM tags")
	if err != nil {
		fmt.Println(err)
		return []string{}
	}
	var repoTags []string
	var tag string
	for tags.Next() {
		fmt.Println()
		tags.Scan(&tag)
		repoTags = append(repoTags, tag)
	}
	return repoTags
}

// GetRepoID retrieves from the database the mapped ID for the repository (inserts if non existing)
func (t *TagsStorage) GetRepoID(repoID string) int64 {
	result, err := t.db.Exec("INSERT INTO repos (githubID) VALUES (?)", repoID)
	if err != nil {
		row := t.db.QueryRow("SELECT id FROM repos WHERE githubID LIKE ?", repoID)
		var id int64
		err = row.Scan(&id)
		if err != nil {
			fmt.Println(err)
			return -1
		}
		return id
	}
	id, _ := result.LastInsertId()
	return id
}

// GetRepoTags retrieves from the database all tags of a specific repository
func (t *TagsStorage) GetRepoTags(repoID int64) []string {
	tags, err := t.db.Query("SELECT t.tag FROM taggedRepos AS tr LEFT JOIN tags AS t ON t.id = tr.tagID WHERE repoID = ?", repoID)
	if err != nil {
		fmt.Println(err)
		return []string{}
	}
	var repoTags []string
	var tag string
	for tags.Next() {
		tags.Scan(&tag)
		repoTags = append(repoTags, tag)
	}
	return repoTags
}

// GetReposByTagPattern Retrieves from the database the repositories which the tags match the mattern
func (t *TagsStorage) GetReposByTagPattern(tagPattern string) []string {
	repos, err := t.db.Query("SELECT DISTINCT repos.githubID FROM repos INNER JOIN taggedRepos AS tr ON tr.repoID = repos.id WHERE tr.tagID IN (SELECT id FROM tags WHERE tag LIKE ?)", tagPattern + "%")
	if err != nil {
		fmt.Println(err)
		return []string{}
	}
	var matchedRepos []string
	var repo string
	for repos.Next() {
		repos.Scan(&repo)
		matchedRepos = append(matchedRepos, repo)
	}
	return matchedRepos
}

// AddRepoTag adds a tag to a repository
func (t *TagsStorage) AddRepoTag(repoID int64, tag string) {
	if _, err := t.db.Exec("INSERT INTO tags (tag) VALUES (?)", tag); err != nil {
		fmt.Println(err)
	}

	if _, err := t.db.Exec("INSERT INTO taggedRepos (repoID, tagID) SELECT ? as repoID, id as tagID FROM tags WHERE tag LIKE ?", repoID, tag); err != nil {
		fmt.Println(err)
	}
}

// RemoveRepoTag removes a tag from the repository
func (t *TagsStorage) RemoveRepoTag(repoID int64, tag string) {
	t.db.Exec("DELETE FROM taggedRepos WHERE repoID = ? AND tagID LIKE (SELECT id FROM tags WHERE tag LIKE ?)", repoID, tag)
}

// ClearRepoTags clear tags associated with the repository
func (t *TagsStorage) ClearRepoTags(repoID int64) {
	t.db.Exec("DELETE FROM taggedRepos WHERE repoID = ?", repoID)
}

// DeleteTag clear tags associated with the repository
func (t *TagsStorage) DeleteTag(tag string) {
	t.db.Exec("DELETE FROM tags WHERE tag LIKE ?", tag)
}

// NewTagsStorage return a database handler
func NewTagsStorage() *TagsStorage {
	return &TagsStorage{db: initDatabase()}
}

func initDatabase() *sql.DB {
    database, _ := sql.Open("sqlite3", "./githubstars.db")
    if _, err := database.Exec("CREATE TABLE IF NOT EXISTS tags (id INTEGER PRIMARY KEY, tag TEXT UNIQUE)"); err != nil {
		log.Fatal(err)
	}
	if _, err := database.Exec("CREATE TABLE IF NOT EXISTS repos (id INTEGER PRIMARY KEY, githubID TEXT UNIQUE)"); err != nil {
		log.Fatal(err)
	}
	if _, err := database.Exec("CREATE TABLE IF NOT EXISTS taggedRepos (tagID INTEGER, repoID INTEGER, UNIQUE(tagID, repoID), FOREIGN KEY(tagID) REFERENCES tags (id) ON DELETE CASCADE, FOREIGN KEY(repoID) REFERENCES repos (id) ON DELETE CASCADE)"); err != nil {
		log.Fatal(err)
	}
    return database
}
