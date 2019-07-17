package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// SearchIssues searches for github issues based on the search terms
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Query failed: %s", resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		resp.Body.Close()
		return nil, err
	}
	result := new(IssuesSearchResult)
	err = json.Unmarshal(data, result)
	if err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return result, nil
}

// SearchIssues2 searches for github issues based on the search terms
func SearchIssues2(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Query failed: %s", resp.Status)
	}

	var result = new(IssuesSearchResult)
	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return result, nil
}
