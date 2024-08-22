// models/word.go
package models

import (
	"yt-nexus-db/database"
)

// FetchAllWords retrieves all words from the dictionary.
func FetchAllWords() ([]string, error) {
	query := `SELECT word FROM dictionary`

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var words []string
	for rows.Next() {
		var word string
		if err := rows.Scan(&word); err != nil {
			return nil, err
		}
		words = append(words, word)
	}

	return words, nil
}
