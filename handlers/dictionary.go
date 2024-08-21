package handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"yt-nexus-db/database"
)

func AddWords(c *gin.Context) {
	var words []string
	if err := c.ShouldBindJSON(&words); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a channel to collect the word IDs
	wordIDs := make(chan int, len(words))
	errChan := make(chan error, len(words))

	// Process each word asynchronously
	for _, word := range words {
		go func(w string) {
			id, err := addWordToDictionary(w)
			if err != nil {
				errChan <- err
			} else {
				wordIDs <- id
			}
		}(word)
	}

	// Collect the results
	var successfulIDs []int
	var errors []error
	for i := 0; i < len(words); i++ {
		select {
		case id := <-wordIDs:
			successfulIDs = append(successfulIDs, id)
		case err := <-errChan:
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		var errorMessages []string
		for _, err := range errors {
			errorMessages = append(errorMessages, err.Error())
		}
		c.JSON(http.StatusInternalServerError, gin.H{"errors": errorMessages})
		return
	}

	c.JSON(http.StatusOK, gin.H{"word_ids": successfulIDs})
}

func addWordToDictionary(word string) (int, error) {
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
