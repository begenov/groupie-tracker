package pkg

import (
	"errors"
	"fmt"
	"search-bar/internal/model"
	"strconv"
	"strings"
)

func Filter(member []string, firstd1, firstd2, creationD1, creationD2, location string) ([]model.Artist, error) {
	// we take apart our api
	art, err := Parsing()
	// checking error
	if err != nil {
		return []model.Artist{}, err
	}
	// we create our variable array struktura
	var result []model.Artist
	// If our creation date is blank
	if creationD1 == "" && creationD2 == "" {
		creationD1 = "1958"
		creationD2 = "2017"
	}
	// we are transporting our creation date on the int
	// We used atoi
	// atoi takes a string and returns int and error
	creat1, err1 := strconv.Atoi(creationD1)
	creat2, err2 := strconv.Atoi(creationD2)
	// checking error
	if err1 != nil && err2 != nil {
		return result, errors.New("ERROR")
	}
	// If our create1 is larger than create2, we return an error
	if creat1 > creat2 {
		return result, errors.New("ERROR")
	}

	// cycle through our structures
	for _, v := range art {
		// here we check the creation date, if it is between creat1 and creat2
		if creat1 <= v.CreationDate && v.CreationDate <= creat2 {
			// in this condition we check if our input is in our structure Relations
			for q := range v.Relations {
				location = strings.Replace(location, ", ", "-", 1)
				//
				if strings.Contains(strings.ToLower(q), strings.ToLower(location)) {
					// If our dick is still zero
					if len(member) == 0 {
						// function contains for our structure so that they do not get repeated
						if contains(result, v) {
							result = append(result, v)
						}
					} else {
						for _, q := range member {
							m, err := strconv.Atoi(q)
							if err != nil {
								return []model.Artist{}, errors.New("error")
							}
							if len(v.Members) == m {
								if contains(result, v) {
									result = append(result, v)
								}
							}
						}
					}
				}
			}
		}
	}
	// check for emptiness of our first data
	if firstd1 != "" && firstd2 != "" {
		res, err := firstdata(result, firstd1, firstd2)
		if err != nil {
			return res, err
		}
		return res, nil
	}
	return result, nil
}

func firstdata(art []model.Artist, s1, s2 string) ([]model.Artist, error) {
	// we create our variable array struktura
	var result []model.Artist
	// create our firstd1 and firstd2 or err1 and err2
	// call the rax function and it takes a string and returns int and error
	firstd1, err1 := rax(s1)
	firstd2, err2 := rax(s2)
	// checking error
	// // If our firstd1 is larger than firstd2, we return an error
	if err1 != nil && err2 != nil {
		return result, err1
	} else if firstd1 > firstd2 {
		fmt.Println("ok")
		return result, errors.New("error")
	}
	// cycle through our structures
	for _, v := range art {
		// call the rax function and it takes a string and returns int and error
		f1, err := rax(v.FirstAlbum)
		// checking error
		if err != nil {
			return result, err
		}
		if firstd1 <= f1 && f1 <= firstd2 {
			if contains(result, v) {
				result = append(result, v)
			}
		}

	}
	return result, nil
}

func rax(s string) (int, error) {
	// Create a line of res
	// Create a line of year
	res := ""
	year := "yyyy"
	// cycle through  our s
	for i := len(s) - 1; i >= 0; i-- {
		// if our rune is not equal to '-'
		if s[i] != '-' {
			// then we will add to our res s[i]
			res = string(s[i]) + res
		} else {
			// If it's all the same, we take a break and he stops
			break
		}
	}
	// check our res if it is not equal to len(year) then we return the error
	if len(res) != len(year) {
		return 0, errors.New("ERROR")
	}
	// We used atoi
	// atoi takes a string and returns int and error
	result, err := strconv.Atoi(res)
	// cheacking error
	if err != nil {
		return 0, err
	}
	// we return the result and nil
	return result, nil
}
