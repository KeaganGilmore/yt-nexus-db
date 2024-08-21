package handlers

import (
	"net/http"
	"yt-nexus-db/models"

	"github.com/gin-gonic/gin"
)

func AddChannel(c *gin.Context) {
	var channel struct {
		ChannelName string `json:"channel_name"`
	}

	if err := c.ShouldBindJSON(&channel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	channelID, err := models.GetOrCreateChannel(channel.ChannelName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"channel_id": channelID})
}
