package pkg

// In the Proverka function we will check the number of artists

func Checker(n int) bool {
	// min artist count
	// You can change this if the artist's maximum increases
	minArtistCount := 0
	// max artist count
	// You can change this if the artist's maximum increases
	maxArtistCount := 53

	// checking artist count
	if n < maxArtistCount && n > minArtistCount {
		return true
	}

	// return bool
	return false
}
