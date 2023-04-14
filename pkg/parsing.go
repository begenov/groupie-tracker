package pkg

import (
	"encoding/json"
	"fmt"
	"search-bar/internal/model"
)

// In the ApiArtist function we retrieve data using "json.unmarshal"
// In the RelationPage function we retrieve data using "json.unmarshal"

func Parsing() ([]model.Artist, error) {
	// In the ParseAPI function we parse the json file
	// ParseAPI takes url and id
	// ParseAPI returns an array of bytes and an error
	bytes, err := ParseAPI("https://groupietrackers.herokuapp.com/api/relation", "")
	// If there is an error
	// then our condition works and it immediately returns an error and an empty array of the structure
	if err != nil {
		fmt.Println(err)
		return []model.Artist{}, err
	}

	// we create our model.Relation structure
	var relation model.Relation

	// and using json.Unmarshal we parse an array of bytes into our structure
	json.Unmarshal(bytes, &relation)

	// In the ParseAPI function we parse the json file
	// ParseAPI takes url and id
	// ParseAPI returns an array of bytes and an error
	b, err := ParseAPI("https://groupietrackers.herokuapp.com/api/artists", "")
	// If there is an error
	// then our condition works and it immediately returns an error and an empty array of the structure
	if err != nil {
		fmt.Println(err)
		return []model.Artist{}, err
	}

	// we create our array structure Model.artist
	var artists []model.Artist

	// and using json.Unmarshal we parse an array of bytes into our structure
	json.Unmarshal(b, &artists)

	// this is our result
	var result []model.Artist

	// we are going through an artists cycle
	// In this cycle we combine the structure of the artist and the relation.
	for i, v := range artists {
		// at the beginning in the v.Relations structure it is empty
		// Here we add from the relation.Index[i].DatesLocations structure
		v.Relations = relation.Index[i].DatesLocations
		// and then add our result
		result = append(result, v)
	}

	// and return our result structure and nil
	return result, nil
}
