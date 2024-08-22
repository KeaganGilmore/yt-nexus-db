// handlers/word.go
package handlers

import (
	"net/http"
	"yt-nexus-db/models"

	"github.com/gin-gonic/gin"
)

// GetAllWords handles the endpoint for retrieving all words.
func GetAllWords(c *gin.Context) {
	words, err := models.FetchAllWords()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"words": words})
}
