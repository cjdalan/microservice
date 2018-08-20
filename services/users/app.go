package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Users")
}

func userLoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "users login")
}

func main() {
	http.HandleFunc("/users", handler)
	http.HandleFunc("/users/login", userLoginHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
