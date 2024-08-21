package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"yt-nexus-db/database"
	"yt-nexus-db/handlers"
)

func main() {
	database.Init()

	r := gin.Default()

	// Set up CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Allows all origins. Adjust as needed for security.
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	ytNexus := r.Group("/yt-nexus")
	{
		ytNexus.POST("/dictionary", handlers.AddWords)
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
