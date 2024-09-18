package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"instagram-api/database"
	"instagram-api/handlers"
	"instagram-api/repository"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()

	userRepository := repository.NewUserRepository(db)
	postRepository := repository.NewPostRepository(db)

	userHandler := handlers.NewUserHandler(userRepository)
	postHandler := handlers.NewPostHandler(postRepository)

	r := mux.NewRouter()

	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}", userHandler.GetUser).Methods("GET")

	r.HandleFunc("/posts", postHandler.CreatePost).Methods("POST")
	r.HandleFunc("/posts/{id:[0-9]+}", postHandler.GetPost).Methods("GET")
	r.HandleFunc("/posts/users/{userId:[0-9]+}", postHandler.ListUserPosts).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Println("Server listening on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
