package models

import (
	"database/sql"
	"yt-nexus-db/database"
)

type Channel struct {
	ID          int    `json:"id"`
	ChannelName string `json:"channel_name"`
}

func GetOrCreateChannel(channelName string) (int, error) {
	var channelID int
	err := database.DB.QueryRow("SELECT id FROM youtube_channels WHERE channel_name = ?", channelName).Scan(&channelID)
	if err == sql.ErrNoRows {
		res, err := database.DB.Exec("INSERT INTO youtube_channels (channel_name) VALUES (?)", channelName)
		if err != nil {
			return 0, err
		}
		channelID64, err := res.LastInsertId()
		if err != nil {
			return 0, err
		}
		channelID = int(channelID64)
	} else if err != nil {
		return 0, err
	}
	return channelID, nil
}
