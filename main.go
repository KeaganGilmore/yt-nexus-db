package main

import (
	"github.com/gin-gonic/gin"
	"yt-nexus-db/database"
	"yt-nexus-db/handlers"
)

func main() {
	database.Init()

	r := gin.Default()

	ytNexus := r.Group("/yt-nexus")
	{
		ytNexus.POST("/dictionary", handlers.AddWord)
		ytNexus.POST("/channel", handlers.AddChannel)
		ytNexus.POST("/video", handlers.AddVideo)
		ytNexus.GET("/channel/:channel_name/common-words", handlers.GetChannelCommonWords)
		ytNexus.GET("/channel/:channel_name/top-videos", handlers.GetChannelTopVideos)
		ytNexus.GET("/channel/:channel_name/keyword/:keyword", handlers.GetVideosWithKeyword)
		ytNexus.GET("/search", handlers.SearchAcrossDB)
		ytNexus.POST("/multi-channel-search", handlers.SearchAcrossChannels)
		ytNexus.POST("/multi-video-search", handlers.SearchAcrossVideos)
	}

	r.Run(":8110")
}
