package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)

	log.Println("Listening on :3001...")
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./website/koob.html")
	return
}
