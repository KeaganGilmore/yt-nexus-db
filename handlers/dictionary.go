package handlers

import (
	"net/http"
	"yt-nexus-db/models"

	"github.com/gin-gonic/gin"
)

func AddWord(c *gin.Context) {
	var word struct {
		Word string `json:"word"`
	}

	if err := c.ShouldBindJSON(&word); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	wordID, err := models.AddWordToDictionary(word.Word)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"word_id": wordID})
}
