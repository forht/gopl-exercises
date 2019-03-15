package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getIssue(user, repo string, id int) error {
	const fs = APIurl + "/repos/%s/%s/issues/%d"
	url := fmt.Sprintf(fs, user, repo, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	resp, err := http.DefaultClient.Do(req)

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("HTTP Errror: %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return err
	}
	fmt.Printf("Title: %s\n", result.Title)
	fmt.Printf("Created at: %v\n", result.CreatedAt)
	fmt.Printf("Body:\n%s\n", result.Body)
	resp.Body.Close()
	return nil
}
