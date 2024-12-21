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

type TagsOutput struct {
	Tags []string `json:"tags"`
}

func getUserTags(userID string) []string {

	// url := "http://account/tags/get"

	// // Create a new HTTP request
	// req, err := http.NewRequest("GET", url, nil)
	// if err != nil {
	// 	log.Fatalf("Error creating request: %v", err)
	// }

	// // Set the x-user-id header
	// req.Header.Set("x-user-id", userID)

	// // Create an HTTP client and send the request
	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	log.Fatalf("Error sending request: %v", err)
	// }
	// defer resp.Body.Close()

	// // Read and print the response body
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalf("Error reading response body: %v", err)
	// }

	// fmt.Println("Response Status:", resp.Status)
	// fmt.Println("Response Body:", string(body))

	// var tagsOutput TagsOutput
	// if err := json.NewDecoder(resp.Body).Decode(&tagsOutput); err != nil {
	// 	fmt.Printf("Error decoding response for user tags %s: %v\n", userID, err)
	// 	return []string{}
	// }

	// tags := tagsOutput.Tags
	// fmt.Printf("For user %s found tags %s\n", userID, tags)

	return []string{"python", "api", "telegram", "developing", "bot"}
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
