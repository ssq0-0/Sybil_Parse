package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
)

const (
	githubToken = "Your Github token"
	repo        = "LayerZero-Labs/sybil-report"
)

type Issue struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Result struct {
	Issue        Issue
	SearchString string
}

func fetchIssues(page int, wg *sync.WaitGroup, results chan<- Result) {
	defer wg.Done()

	url := fmt.Sprintf("https://api.github.com/repos/%s/issues?state=all&page=%d&per_page=30", repo, page)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating request: %v\n", err)
		return
	}
	req.Header.Set("Authorization", "token "+githubToken)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching page %d: %v\n", page, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Failed to fetch issues for page %d: %s\n", page, resp.Status)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading response for page %d: %v\n", page, err)
		return
	}

	var issues []Issue
	if err := json.Unmarshal(body, &issues); err != nil {
		fmt.Fprintf(os.Stderr, "Error unmarshaling response for page %d: %v\n", page, err)
		return
	}

	searchStrings := []string{
		"you_wallet_1",
		"you_wallet_2",
		"you_wallet_3",
		"you_wallet_4",
	}

	for _, issue := range issues {
		body := strings.ToLower(issue.Body)
		title := strings.ToLower(issue.Title)
		for _, searchString := range searchStrings {
			searchString = strings.ToLower(searchString)
			if strings.Contains(title, searchString) || strings.Contains(body, searchString) {
				results <- Result{Issue: issue, SearchString: searchString}
				break
			}
		}
	}
}

func main() {
	var totalPages int
	totalPages = 150

	foundIssues := make(map[int]Issue)
	results := make(chan Result, 100)
	var wg sync.WaitGroup

	concurrency := 15
	sem := make(chan struct{}, concurrency)

	for page := 1; page <= totalPages; page++ {
		wg.Add(1)
		sem <- struct{}{}
		go func(page int) {
			defer func() { <-sem }()
			fmt.Printf("Fetching page %d...\n", page)
			fetchIssues(page, &wg, results)
		}(page)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		issue := result.Issue
		if _, exists := foundIssues[issue.Number]; !exists {
			foundIssues[issue.Number] = issue
			fmt.Printf("Found in Issue #%d\n", issue.Number)
			fmt.Printf("Matching string: %s\n", result.SearchString)
			fmt.Printf("Title: %s\n", issue.Title)
			fmt.Println("---------------------------------")
		}
	}

	fmt.Println("Processing complete.")
}
