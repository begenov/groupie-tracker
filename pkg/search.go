package pkg

import (
	"search-bar/internal/model"
	"strconv"
	"strings"
)

// function Search is our search engine
// The function search engine accepts our login page, and all artist data
func Search(s string, res []model.Artist) []model.Artist {
	// we create the result of the structure
	var result []model.Artist

	// Going through all the artists through the cycle
	for _, v := range res {

		// strings.Contains Function in Golang is used to check the given letters present in the given string or not.
		// If the letter is present in the given string, then it will return true, otherwise, return false.
		// To convert string to lowercase in Go programming, callstrings.
		// ToLower() function and pass the string as argument to this function.

		// in this condition we check if our input is in our structure name
		if strings.Contains(strings.ToLower(v.Name), strings.ToLower(s)) {
			// function contains for our structure so that they do not get repeated
			if contains(result, v) {
				result = append(result, v)
			}
		}

		// in this condition we check if our input is in our structure Members
		for _, j := range v.Members {
			if strings.Contains(strings.ToLower(j), strings.ToLower(s)) {
				// function contains for our structure so that they do not get repeated
				if contains(result, v) {
					result = append(result, v)
				}
			}
		}
		// in this condition we check if our input is in our structure CreationDate
		if strings.Contains(strconv.Itoa(v.CreationDate), s) {
			// function contains for our structure so that they do not get repeated
			if contains(result, v) {
				result = append(result, v)
			}
		}
		// in this condition we check if our input is in our structure FirstAlbum
		if strings.Contains(v.FirstAlbum, s) {
			// function contains for our structure so that they do not get repeated
			if contains(result, v) {
				result = append(result, v)
			}
		}
		// in this condition we check if our input is in our structure Relations
		for i, l := range v.Relations {
			if strings.Contains(strings.ToLower(i), strings.ToLower(s)) {
				// function contains for our structure so that they do not get repeated
				if contains(result, v) {
					result = append(result, v)
				}
			}
			for q := 0; q < len(l); q++ {
				// function contains for our structure so that they do not get repeated
				if strings.Contains(strings.ToLower(l[q]), strings.ToLower(s)) {
					if contains(result, v) {
						result = append(result, v)
					}
				}
			}
		}
	}

	// return our result structre
	return result
}

// function contains for our structure so that they do not get repeated
func contains(arr []model.Artist, str model.Artist) bool {
	for _, v := range arr {
		if v.Name == str.Name {
			return false
		}
	}
	return true
}
