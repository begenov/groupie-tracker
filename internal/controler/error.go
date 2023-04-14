package controler

import (
	"html/template"
	"net/http"
	"search-bar/internal/model"
)

//  In the error function take the error as int and ResponseWriter

func Error(status int, w http.ResponseWriter) {
	// we are creating our model structure here
	var errors model.Errors

	// http.StatusText takes our status and returns it as a string
	errors.ErrorName = http.StatusText(status)

	// WriteHeader sends an HTTP response header with the provided
	// status code.
	w.WriteHeader(status)
	errors.Num = status
	temp, err := template.ParseFiles("templates/error.html")
	//  If tempalte.parsefile cannot find the html file
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError) // then we will use http.Error to cause an error
		return
	}
	if err := temp.Execute(w, errors); err != nil {
		Error(http.StatusInternalServerError, w)
	}
}
