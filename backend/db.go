package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type ArticleResponse struct {
	ID   int      `json:"id"`
	Tags []string `json:"tags"`
}

type TagsOutput struct {
	Tags []string `json:"tags"`
}

type Config struct {
	IDs []int `yaml:"ids"`
}

func GetUserTags(userID string) []string {

	url := fmt.Sprintf("http://%s/tags/get", AppConfig.PersonalAccountHostPort)
	log.Println("Generated URL for personal account:", url)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Set the x-user-id header
	req.Header.Set("x-user-id", userID)

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read and print the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body))

	var tagsOutput TagsOutput
	if err := json.NewDecoder(resp.Body).Decode(&tagsOutput); err != nil {
		fmt.Printf("Error decoding response for user tags %s: %v\n", userID, err)
		return []string{}
	}

	tags := tagsOutput.Tags
	fmt.Printf("For user %s found tags %s\n", userID, tags)

	// Define themes and their associated tags
	// log.Printf("User id: %s", userID)
	// themes := map[string][]string{
	// 	"программирование": {"питон", "го", "java", "джаваскрипт", "си++"},
	// 	"api":              {"rest", "graphql", "soap", "json", "xml"},
	// 	"мессенджеры":      {"телеграм", "слак", "дискорд", "ватсап", "сигнал"},
	// 	"разработка":       {"фронтенд", "бэкенд", "фулстек", "девопс", "облако"},
	// 	"автоматизация":    {"бот", "скриптинг", "ci/cd", "тестирование", "развертывание"},
	// 	"бизнес":           {"маркетинг", "финансы", "продажи", "стратегия", "кадры"},
	// }

	// // Get a list of theme names
	// themeNames := make([]string, 0, len(themes))
	// for theme := range themes {
	// 	themeNames = append(themeNames, theme)
	// }

	// // Select a random theme
	// selectedTheme := themeNames[rand.Intn(len(themeNames))]
	// log.Printf("selectedTheme : %s", selectedTheme)
	// // Return the tags from the selected theme

	return tags
}

// GetArticles reads the articles from a CSV file and returns a map of article IDs to their tags.
func GetArticles() map[string][]string {
	articles := make(map[string][]string)

	file, err := os.Open("/data/articles.csv")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read the header line to skip it
	_, err = reader.Read()
	if err != nil {
		log.Fatalf("failed to read header: %s", err)
	}

	// Read each record from the CSV
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatalf("failed to read record: %s", err)
		}

		// Assuming the CSV columns are: id, name, text, complexity, reading_time, tags
		id := record[0]
		tags := record[5]

		log.Println("Tags: ", tags)

		var tagList []string
		err = json.Unmarshal([]byte(tags), &tagList)
		if err != nil {
			log.Printf("failed to parse tags for article %s: %s", id, err)
			continue
		}

		// Add to the map
		articles[id] = tagList
	}
	log.Printf("Total number of articles: %d", len(articles))
	return articles
}

// Function to get articles and their tags from api {id1: [tag1, tag2, tag3]}
func GetArticlesSlow() map[string][]string {
	articles := make(map[string][]string)

	file, err := os.Open("/data/ids.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Each line is an ID
		articleID := scanner.Text()
		fmt.Println("ID:", articleID)

		url := fmt.Sprintf("http://%s/api/v1/article/%s", AppConfig.RecsDbHostPort, articleID)

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

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	return articles
}
