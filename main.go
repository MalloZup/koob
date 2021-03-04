package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("website")))
	http.HandleFunc("/bookmarks", bookmarks)
	log.Println("Listening on :3001...")
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatal(err)
	}
}

type Bookmarks struct {
	Name string `json:"Name"`
	Tag  string `json:"Tag"`
}

func bookmarks(w http.ResponseWriter, r *http.Request) {

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	var t Bookmarks

	fmt.Println("Entered")
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("error")
		return
	}

	fmt.Println("Bookmark called")
	fmt.Println(t)
	// Do something with the Person struct...
	fmt.Fprintf(w, "Person: %+v", t)

}
