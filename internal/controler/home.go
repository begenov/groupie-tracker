package controler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// In the home function we handle the path "/"

func Home(w http.ResponseWriter, r *http.Request) {
	// we are checking our way
	if r.URL.Path != "/" {
		Error(http.StatusNotFound, w)
		return
	}

	// we are checking our method
	if r.Method != http.MethodGet {
		Error(http.StatusMethodNotAllowed, w)
		return
	}

	// template.ParseFiles we parse to an html file
	// parsefiles returns *tempate and an error
	tmpl, err := template.ParseFiles("templates/index.html")
	// error-checking
	if err != nil {
		fmt.Println("Error not found index.html")
		return
	}

	// Here we do executeTemplate
	// executeTemplate takes the ResposeWriter argument, our html file and any file
	// and returns error
	if err := tmpl.ExecuteTemplate(w, "index.html", nil); err != nil {
		log.Fatalf("Error in index.html: %s", err)
	}
	// error-checking
}
