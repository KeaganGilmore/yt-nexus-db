package handlers

import (
	"net/http"
	"yt-nexus-db/models"

	"github.com/gin-gonic/gin"
)

// GetChannelCommonWords retrieves the most common words used in a specific channel's videos.
func GetChannelCommonWords(c *gin.Context) {
	channelName := c.Param("channel_name")

	commonWords, err := models.FetchCommonWordsByChannel(channelName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"common_words": commonWords})
}

// GetChannelTopVideos retrieves the top videos from a specific channel based on some criteria (e.g., word count, etc.).
func GetChannelTopVideos(c *gin.Context) {
	channelName := c.Param("channel_name")

	topVideos, err := models.FetchTopVideosByChannel(channelName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"top_videos": topVideos})
}

// GetVideosWithKeyword retrieves videos from a specific channel that contain a particular keyword.
func GetVideosWithKeyword(c *gin.Context) {
	channelName := c.Param("channel_name")
	keyword := c.Param("keyword")

	videoWordCounts, err := models.FetchVideosByKeyword(channelName, keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"videos": videoWordCounts})
}
func AddVideo(c *gin.Context) {
	var video struct {
		ChannelID  int         `json:"channel_id"`
		VideoID    string      `json:"video_id"`
		WordCounts map[int]int `json:"word_counts"`
	}

	if err := c.ShouldBindJSON(&video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.InsertVideo(video.ChannelID, video.VideoID, video.WordCounts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Video added successfully!"})
}

// SearchAcrossDB searches for a keyword across all channels and videos in the database.
func SearchAcrossDB(c *gin.Context) {
	keyword := c.Query("keyword")
	videoWordCounts, err := models.FetchVideosWithKeywordAcrossDB(keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"videos": videoWordCounts})
}

// SearchAcrossChannels searches for a keyword across multiple channels.
func SearchAcrossChannels(c *gin.Context) {
	var channelNames []string
	if err := c.ShouldBindJSON(&channelNames); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	keyword := c.Query("keyword")
	videoWordCounts, err := models.FetchVideosWithKeywordAcrossChannels(channelNames, keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"videos": videoWordCounts})
}

// SearchAcrossVideos searches for a keyword across multiple videos.
func SearchAcrossVideos(c *gin.Context) {
	var videoIDs []string
	if err := c.ShouldBindJSON(&videoIDs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	keyword := c.Query("keyword")
	videoWordCounts, err := models.FetchVideosWithKeywordAcrossVideos(videoIDs, keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"videos": videoWordCounts})
}
