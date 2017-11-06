package main

import (
	"fmt"
	"log"
	"net/http"
)

// func messageHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to Go Web Development")
// }
func messageHandler(message string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, message)
	}
}

func main() {
	http.HandleFunc("/welcome", messageHandler("Welcome to Go Web Development"))
	http.HandleFunc("/message", messageHandler("net/http is awesome"))

	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
