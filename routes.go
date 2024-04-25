package main

import (
	"encoding/json"
	"golang-rest-api/entity"
	"golang-rest-api/repository"
	"math/rand"
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

var (
	repo repository.PostRepository = repository.NewPostRepository()
)

func getPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error getting the posts"}`))
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}

func addPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var post entity.Post
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error unmarshalling data"}`))
		return
	}
	post.ID = rand.Int63()
	repo.Save(&post)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(post)

}
