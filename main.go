package main

import (
	"encoding/json"
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
	Name string
	Tag  string
}

func bookmarks(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var t Bookmarks
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	log.Println(t)
	log.Println(t.Name)
}
