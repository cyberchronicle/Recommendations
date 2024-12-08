package main

func getUserTags(userId string) []string {
	return []string{"a", "b", "c", "d", "e"}
}

// return dict: key - article_id(string), value - tags(string array)
func getArticles() map[string][]string {
	return map[string][]string{
		"aid1": {"a", "b", "c"},
		"aid2": {"c", "d", "e"},
	}
}
