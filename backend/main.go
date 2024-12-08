package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/recommendations/relevant", relevantRecommendations)

	http.ListenAndServe(":8383", nil)
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
	articles := getArticles([]string{"858174", "859996", "859692"})

	suggestedArticles := suggestArticles(userTags, articles)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(suggestedArticles[min(offset, len(suggestedArticles)):])


	// just for testing by trigger
	as := map[string]string{
        "1": "что такое осень это небо",
        "2": "у меня зима на сердце, на душе вьюга",
    }
	
	for articleID, text := range as {
		if len(text) > 0 {
			tags := processText(text)

			log.Printf("Article ID: %s, Tags: %v\n", articleID, tags)
		}
	}
}
