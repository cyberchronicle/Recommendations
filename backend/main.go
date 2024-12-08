package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/recommendations/relevant", relevantRecommendations)

	http.ListenAndServe(":8383", nil)
}

func relevantRecommendations(w http.ResponseWriter, r *http.Request) {
	idRaw := r.URL.Query()["id"]
	offsetRaw := r.URL.Query()["offset"]

	if len(idRaw) == 0 || idRaw[0] == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("query parameter 'id' is required - user id"))
		return
	}
	id := idRaw[0]
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

	userTags := getUserTags(id)
	articles := getArticles()

	suggestedArticles := suggestArticles(userTags, articles)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(suggestedArticles[min(offset, len(suggestedArticles)):])
}
