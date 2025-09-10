// Package main demonstrates how to use the go-meta-extractor library
// to extract meta tags from a web page
package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	metaextractor "github.com/mrz1836/go-meta-extractor"
)

func main() {
	// Set a client
	client := &http.Client{Timeout: 10 * time.Second}

	// Start the request
	req, err := http.NewRequestWithContext(
		context.Background(), http.MethodGet, "https://mrz1818.com", nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Fire the request
	var resp *http.Response
	if resp, err = client.Do(req); err != nil {
		log.Fatal(err)
	}

	// Close the body
	defer func() {
		_ = resp.Body.Close()
	}()

	// Extract the meta tags
	tags := metaextractor.Extract(resp.Body)

	// Show the tags we found:
	jsonData, err := json.Marshal(tags)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(jsonData))
}
