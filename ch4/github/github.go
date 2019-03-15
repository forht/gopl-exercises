// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 110.
//!+

// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.

package main

import "time"

const APIurl = "https://api.github.com"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int       `json:",omitempty"`
	HTMLURL   string    `json:"html_url,omitempty"`
	Title     string    `json:",omitempty"`
	State     string    `json:",omitempty"`
	User      *User     `json:",omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Body      string    `json:",omitempty"` // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

//!-
