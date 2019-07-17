package github

import (
	"time"
)

// IssuesURL is the public github url for the issue search api
const IssuesURL = "https://api.github.com/search/issues"

// IssuesSearchResult has the important info from a github issues search response
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// Issue represents one github issue
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

// User represents a user
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
