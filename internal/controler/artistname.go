package controler

import (
	"fmt"
	"html/template"
	"net/http"
	"search-bar/pkg"
	"strconv"
)

// In the home function we handle the path "/artistname"
func Artistname(w http.ResponseWriter, r *http.Request) {
	// we are checking our method
	if r.Method != http.MethodGet {
		Error(http.StatusMethodNotAllowed, w)
		return
	}

	// we are checking our way
	if r.URL.Path != "/artistname" {
		Error(http.StatusNotFound, w)
		return
	}

	// template.ParseFiles we parse to an html file
	// parsefiles returns *tempate and an error
	temp, err := template.ParseFiles("templates/artistname.html")
	if err != nil {
		Error(http.StatusInternalServerError, w)
		return
	}

	// this is our artist id
	id := r.FormValue("id")

	// Our default string type identifier
	// we use strconv.atoi to make it an int type
	//  strconv.atoi takes a string type and returns int and error
	num, err := strconv.Atoi(id)
	// here we check for an error
	// If it can't convert our string to int type, it panic
	if err != nil {
		Error(http.StatusNotFound, w)
		return
	}

	// The Checker function we check our id to see if the artifact is such an id
	if !pkg.Checker(num) {
		Error(http.StatusNotFound, w)
		return
	}

	// In the parsing function we parse our JSON file
	// The parsing function returns the array structure model.artist and error
	result, err := pkg.Parsing()
	// If it can't parse our json, it returns an error
	if err != nil {
		Error(http.StatusInternalServerError, w)
		return
	}

	// Here we do executeTemplate
	// executeTemplate takes the ResposeWriter argument, our html file and any file
	// and returns error
	if err := temp.ExecuteTemplate(w, "artistname.html", result[num-1]); err != nil {
		fmt.Println(err.Error())
		Error(http.StatusInternalServerError, w)
		return
	}
	// error-checking
}
