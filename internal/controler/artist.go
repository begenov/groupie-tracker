package controler

import (
	"fmt"
	"html/template"
	"net/http"

	"search-bar/internal/model"
	"search-bar/pkg"
)

// In the home function we handle the path "/artist"
func Artistpage(w http.ResponseWriter, r *http.Request) {
	// we are checking our method
	if r.Method != http.MethodGet {
		Error(http.StatusMethodNotAllowed, w)
		return
	}

	// we are checking our way
	if r.URL.Path != "/artist" {
		fmt.Println("asd")
		Error(http.StatusNotFound, w)
		return
	}

	// we create our result type structure
	// model.Artist it is an array structure
	var result []model.Artist

	// template.ParseFiles we parse to an html file
	// parsefiles returns *tempate and an error
	temp, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		fmt.Println("Hello artist")
		Error(http.StatusInternalServerError, w)
		return
	}
	// variable input is my search bar
	input := r.FormValue("input1")
	// If my input is not empty then our search bar works
	if input != "" {
		// we parse our api
		resinput, err := pkg.Parsing()
		// checkimg error
		if err != nil {
			Error(http.StatusInternalServerError, w)
			return
		}
		// we will send our resinput and our input to the search bar
		result = pkg.Search(input, resinput)
		// Here we do executeTemplate
		// executeTemplate takes the ResposeWriter argument, our html file and any file
		// and returns error
		// error-checking
		if err := temp.ExecuteTemplate(w, "artist.html", result); err != nil {
			fmt.Println(err)
		}
	} else {
		// If my input is blank, our search bar does not work

		// we parse our api
		result, err = pkg.Parsing()

		// check in error
		if err != nil {
			Error(http.StatusInternalServerError, w)
			return
		}
	}

	// Here we create the filters variable
	// Here's our list of artists
	member := r.Form["member"]

	// the first data 1 is our minimum date
	firstd1 := r.FormValue("min")
	// the first data 2 is our maximum date
	firstd2 := r.FormValue("max")

	// the creation data 1 is our minimum date
	creationD1 := r.FormValue("min-2")
	// the creation data 2 is our maximum date
	creationD2 := r.FormValue("max-2")
	// This is our artist location
	location := r.FormValue("locations")

	// all of our variables we will parse into our filter function
	resultfilter, err := pkg.Filter(member, firstd1, firstd2, creationD1, creationD2, location)
	// checking error
	if err != nil {
		Error(http.StatusBadRequest, w)
		return
	}
	// If our result filter is not zero, our condition works
	if len(resultfilter) != 0 {
		// Here we do executeTemplate
		// executeTemplate takes the ResposeWriter argument, our html file and any file
		// and returns error
		// error-checking
		if err := temp.ExecuteTemplate(w, "artist.html", resultfilter); err != nil {
			fmt.Println(err)
		}
		return
	} else {
		// Here we do executeTemplate
		// executeTemplate takes the ResposeWriter argument, our html file and any file
		// and returns error
		// error-checking
		if err := temp.ExecuteTemplate(w, "artist.html", result); err != nil {
			fmt.Println(err)
		}
	}
}
