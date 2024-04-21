package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// func getPosts(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-type", "application/json")

// 	result, err := json.Marshal(posts)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(`{"error": "Error marshalling the posts array"}`))
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(result)

// }

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	posts []Post
)

func init() {
	posts = []Post{Post{Id: 1, Title: "Title 1", Text: "Text 1"}}
	// posts = []Post{{1, "title 1", "text 1"}}
}

func getPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	result, err := json.Marshal(posts)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error marshalling data"}`))
	}
	response.WriteHeader(http.StatusOK)
	response.Write(result)
}

func addPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var post Post
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error unmarshalling data"}`))
		return
	}
	post.Id = len(posts) + 1
	posts = append(posts, post)
	response.WriteHeader(http.StatusOK)
	result, err := json.Marshal(post)
	response.Write(result)

	fmt.Print(post.Id)
}
