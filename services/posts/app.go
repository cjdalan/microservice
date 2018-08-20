package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Posts")
}

func main() {
	http.HandleFunc("/posts", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
