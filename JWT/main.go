package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/welcome", Welcome)
	http.HandleFunc("/refresh", Refresh)
	http.HandleFunc("/logout", Logout)

	//Start server on Port 3000
	log.Fatal(http.ListenAndServe(":8000", nil))

}
