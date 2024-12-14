package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ArticleResponse struct {
	ID   int      `json:"id"`
	Tags []string `json:"tags"`
}

func getUserTags(userId string) []string {
	return []string{"a", "b", "c", "d", "e"}
}

// Function to get articles and their tags
func getArticles(articleIDs []string) map[string][]string {
	articles := make(map[string][]string)

	for _, articleID := range articleIDs {
		// Construct the URL for the article
		url := fmt.Sprintf("http://fastapi:8000/api/v1/article/%s", articleID)

		// Send the GET request to the article service
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error making request for article %s: %v\n", articleID, err)
			continue
		}
		defer resp.Body.Close()

		// Check if the response status is OK
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Received non-OK response for article %s: %s\n", articleID, resp.Status)
		}
		if resp.StatusCode != 201 {
			continue
		}

		// Decode the response body
		var articleResponse ArticleResponse
		if err := json.NewDecoder(resp.Body).Decode(&articleResponse); err != nil {
			fmt.Printf("Error decoding response for article %s: %v\n", articleID, err)
			continue
		}

		// Add the tags to the map
		articles[articleID] = articleResponse.Tags
		fmt.Printf("For article %s found tags %s\n", articleID, articleResponse.Tags)
		
	}
	return articles
}