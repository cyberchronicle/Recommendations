package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	err := LoadConfig("/configs/service.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Use the configuration
	log.Println("ML Host Port:", AppConfig.MLHostPort)
	log.Println("Backend Host Port:", AppConfig.BackendHostPort)
	log.Println("Personal Account Host Port:", AppConfig.PersonalAccountHostPort)
	log.Println("Recs Db Host Port:", AppConfig.RecsDbHostPort)
	log.Println("Scrapper Host Port:", AppConfig.ScrapperHostPort)

	http.HandleFunc("/recommendations/relevant", relevantRecommendations)

	parts := strings.SplitN(AppConfig.BackendHostPort, ":", 2)
	var host, port string

	if len(parts) == 2 {
		host = parts[0]
		port = parts[1]

		fmt.Println("Backend Host:", host)
		fmt.Println("Backend Port:", port)
	} else {
		host = "localhost"
		port = "8383"
		fmt.Println("The input string does not contain the ':' character. Listen and serve at port:", port)
	}

	addr := fmt.Sprintf(":%s", port)
	log.Printf("Starting server on %s", addr)
	http.ListenAndServe(addr, nil)
}

func relevantRecommendations(w http.ResponseWriter, r *http.Request) {
	// example: http://localhost:8383/recommendations/relevant?offset=1
	log.Println("Received request:", r.URL)

	// Extract user ID from the headers
	userID := r.Header.Get("X-User-Id")
	if userID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("header 'X-User-Id' is required"))
		return
	}
	log.Printf("User ID: %s", userID)

	// Extract offset from query parameters
	offsetRaw := r.URL.Query()["offset"]
	offset := 0
	if len(offsetRaw) != 0 {
		var err error
		offset, err = strconv.Atoi(offsetRaw[0])
		if err != nil && offsetRaw[0] != "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte("query parameter 'offset' incorrect - number expected"))
			return
		}
	}
	log.Printf("Offset: %d", offset)

	userTags := GetUserTags(userID)

	suggestedArticles := suggestArticles(userTags)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(suggestedArticles[min(offset, len(suggestedArticles)):])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
