package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yt-nexus-db/models"
)

func AddWords(c *gin.Context) {
	var words []string
	if err := c.ShouldBindJSON(&words); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	wordIDs := make(chan int, len(words))
	errChan := make(chan error, len(words))

	for _, word := range words {
		go func(w string) {
			id, err := models.AddWordToDictionary(w) // Use the exported function here
			if err != nil {
				errChan <- err
			} else {
				wordIDs <- id
			}
		}(word)
	}

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
