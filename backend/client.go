package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func suggestArticles(userTags []string, arcticles map[string][]string) []string {
	return []string{"1", "2", "3", "4", "5"}
}

type TextRequest struct {
	Text string `json:"text"`
}

type TagsResponse struct {
	Tags []string `json:"tags"`
}

func processText(text string) []string {
	ml_text_process_url := "http://ml:8001/text/process"

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
