package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func setAuthorization(req *http.Request, in *bufio.Reader) error {
	fmt.Print("Insert username: ")
	username, err := in.ReadString('\n')
	fmt.Print("Insert password: ")
	password, err := in.ReadString('\n')
	if err != nil {
		return err
	}
	username = strings.TrimRight(username, "\n")
	password = strings.TrimRight(password, "\n")
	handle := username + ":" + password
	encoded := base64.StdEncoding.EncodeToString([]byte(handle))
	req.Header.Set("Authorization", "Basic "+encoded)
	return nil
}

func createIssue(user, repo string) error {
	type buildIssue struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	in := bufio.NewReader(os.Stdin)
	fmt.Printf("Insert title of the issue: ")
	title, err := in.ReadString('\n')
	fmt.Printf("Insert body of the issue: ")
	body, err := in.ReadString('\n')
	if err != nil {
		return err
	}
	issue := buildIssue{Title: title, Body: body}
	bs, err := json.Marshal(issue)
	if err != nil {
		return err
	}
	url := fmt.Sprintf(APIurl+"/repos/%s/%s/issues", user, repo)

	req, err := http.NewRequest("POST", url, bytes.NewReader(bs))
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")

	if err = setAuthorization(req, in); err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusCreated {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
		resp.Body.Close()
		return fmt.Errorf("HTTP Error: %s", resp.Status)
	}
	resp.Body.Close()
	return nil
}
