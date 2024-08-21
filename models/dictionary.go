package models

import (
	"database/sql"
	"yt-nexus-db/database"
)

func AddWordToDictionary(word string) (int, error) {
	var wordID int
	err := database.DB.QueryRow("SELECT id FROM dictionary WHERE word = ?", word).Scan(&wordID)
	if err == sql.ErrNoRows {
		res, err := database.DB.Exec("INSERT INTO dictionary (word) VALUES (?)", word)
		if err != nil {
			return 0, err
		}
		id, err := res.LastInsertId()
		if err != nil {
			return 0, err
		}
		wordID = int(id)
	} else if err != nil {
		return 0, err
	}
	return wordID, nil
}
