package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type TextRequest struct {
	Text string `json:"text"`
}

type TagsResponse struct {
	Tags []string `json:"tags"`
}

// SuggestRequest represents the request payload for the /suggest endpoint
type SuggestRequest struct {
	UserTags []string `json:"user_tags"`
}

// SuggestResponse represents the response payload from the /suggest endpoint
type SuggestResponse struct {
	IDs []string `json:"ids"`
}

func suggestArticles(userTags []string) []string {
	var articleIDs []string
	mlRelevantURL := fmt.Sprintf("http://%s/suggest", AppConfig.MLHostPort)

	// Create the request payload
	requestPayload := SuggestRequest{
		UserTags: userTags,
	}

	// Marshal the request payload into JSON
	jsonData, err := json.Marshal(requestPayload)
	if err != nil {
		log.Printf("Error marshalling request payload: %v", err)
		return articleIDs
	}

	// Create a new HTTP POST request
	resp, err := http.Post(mlRelevantURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error making POST request: %v", err)
		return articleIDs
	}
	defer resp.Body.Close()

	// Check for a successful response
	if resp.StatusCode != http.StatusOK {
		log.Printf("Received non-OK response: %d", resp.StatusCode)
		return articleIDs
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return articleIDs
	}

	// Unmarshal the response JSON into the SuggestResponse struct
	var response SuggestResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Printf("Error unmarshalling response: %v", err)
		return articleIDs
	}

	// Return the list of article IDs
	return response.IDs
}

func processText(text string) []string {
	ml_text_process_url := fmt.Sprintf("http://%s/text/process", AppConfig.MLHostPort)

	requestBody, err := json.Marshal(TextRequest{Text: text})
	if err != nil {
		fmt.Println("Error marshalling request:", err)
		return nil
	}

	resp, err := http.Post(ml_text_process_url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Received non-OK response:", resp.Status)
		return nil
	}

	var tagsResponse TagsResponse
	if err := json.NewDecoder(resp.Body).Decode(&tagsResponse); err != nil {
		fmt.Println("Error decoding response:", err)
		return nil
	}

	return tagsResponse.Tags
}
