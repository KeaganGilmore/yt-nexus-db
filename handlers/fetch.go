package handlers

import (
	"net/http"
	"yt-nexus-db/models"

	"github.com/gin-gonic/gin"
)

func GetAllData(c *gin.Context) {
	words, err := models.FetchAllWords()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch words"})
		return
	}

	channels, err := models.FetchAllChannels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch channels"})
		return
	}

	videoIDs, err := models.FetchAllVideoIDs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch video IDs"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"words":    words,
		"channels": channels,
		"videos":   videoIDs,
	})
}
