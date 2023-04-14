package main

import (
	"log"
	"net/http"

	"search-bar/internal/controler"
)

// program run
func main() {
	// we are creating a new server mux here
	mux := http.NewServeMux()
	log.Println("Running in a web application: http://localhost:8081/")
	mux.HandleFunc("/", controler.Home)

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	mux.HandleFunc("/artist", controler.Artistpage)

	mux.HandleFunc("/artistname", controler.Artistname)
	if err := http.ListenAndServe(":8081", mux); err != nil {
		log.Fatalln(err)
	}
}
