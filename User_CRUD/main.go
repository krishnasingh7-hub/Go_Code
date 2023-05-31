package main

import (
	"log"
	"net/http"
)

func main() {
	// Initialize the database connection
	err := InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// Register the HTTP handlers
	http.HandleFunc("/users", ListUsersHandler)
	http.HandleFunc("/users/create", CreateUserHandler)
	http.HandleFunc("/users/update", UpdateUserHandler)
	http.HandleFunc("/users/delete", DeleteUserHandler)

	// Start the HTTP server
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
