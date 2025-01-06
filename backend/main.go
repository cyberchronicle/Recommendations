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
	http.ListenAndServe(addr, nil)
}

func relevantRecommendations(w http.ResponseWriter, r *http.Request) {
	// example: http://localhost:8383/recommendations/relevant?id=1&offset=1
	log.Println("Received request:", r.URL)

	idRaw := r.URL.Query()["id"]
	offsetRaw := r.URL.Query()["offset"]

	log.Println("idRaw:", idRaw)
	log.Println("offsetRaw:", offsetRaw)

	if len(idRaw) == 0 || idRaw[0] == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("query parameter 'id' is required - user id"))
		return
	}
	id := idRaw[0]
	log.Printf("User ID: %s", id)

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

	userTags := getUserTags(id)
	// articles := getArticles()
	articles := map[string][]string{}

	suggestedArticles := suggestArticles(userTags, articles)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(suggestedArticles[min(offset, len(suggestedArticles)):])

}
