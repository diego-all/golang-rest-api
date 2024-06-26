package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	const port string = ":8000"

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "up and running...")
	})

	//router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", addPost).Methods("POST")
	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))

}

// Trevor se engancha facil
// 	srv := &http.Server{
// 		Addr:    fmt.Sprintf(":%d", app.config.port),
// 		Handler: app.routes(),
// 	}
