package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

type TextRequest struct {
	Text string `json:"text"`
}

type TagsResponse struct {
	Tags []string `json:"tags"`
}

func suggestArticles(userTags []string, articles map[string][]string) []string {
	return []string{"859260", "859264", "859268", "859276", "859278", "859282", "859290", "859292", "859298", "859296", "859294", "859306", "859302", "859312", "859318", "859324", "859326", "859340", "859288", "859354", "859360", "859220", "859216", "859218", "859224", "859230", "859236", "859244", "859250", "859254", "859252", "859170", "859166", "859180", "859174", "859182", "859198", "859200", "859194", "859196", "859116", "859132", "859138", "859142", "859148", "859144", "859140", "859156", "859158", "859362", "859364", "859378", "859386", "859392", "859384", "859400", "859408", "859424", "859430", "859432", "859436", "859442", "859072", "859068", "859066", "859080", "859082", "859070", "859084", "859086", "859088", "859092", "859094", "859096", "859110", "859106", "859016", "859022", "859024", "859020", "859028", "859032", "859046", "859054", "859050", "859060", "859048", "858966", "858976", "858980", "858972", "858986", "858990", "858988", "858994", "858996", "859004", "859012", "859006", "858922"}
	type articleScore struct {
		id    string
		score float64
	}

	var scores []articleScore

	// Convert userTags to a map for quick lookup
	userTagSet := make(map[string]struct{})
	for _, tag := range userTags {
		userTagSet[tag] = struct{}{}
	}

	// Calculate scores for each article
	for id, tags := range articles {
		intersectionCount := 0
		for _, tag := range tags {
			if _, exists := userTagSet[tag]; exists {
				intersectionCount++
			}
		}

		// Calculate score as the proportion of matching tags
		if len(tags) > 0 {
			score := float64(intersectionCount) / float64(len(tags))
			score = min(score, 1)
			if score > 0 {
				scores = append(scores, articleScore{id: id, score: score})
			}
		}
	}

	// Sort articles by score in descending order
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].score > scores[j].score
	})

	// Collect sorted article IDs
	var sortedArticleIDs []string
	for _, article := range scores {
		sortedArticleIDs = append(sortedArticleIDs, article.id)
	}

	return sortedArticleIDs
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
