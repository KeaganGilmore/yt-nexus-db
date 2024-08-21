package models

import (
	"strings"
	"yt-nexus-db/database"
)

// FetchCommonWordsByChannel retrieves the most common words used in a specific channel's videos.
func FetchCommonWordsByChannel(channelName string) (map[string]int, error) {
	query := `
		SELECT d.word, SUM(wc.count) AS total_count
		FROM word_counts wc
		JOIN video_details vd ON wc.video_id = vd.id
		JOIN dictionary d ON wc.word_id = d.id
		JOIN youtube_channels yc ON vd.channel_id = yc.id
		WHERE yc.channel_name = ?
		GROUP BY d.word
		ORDER BY total_count DESC
		LIMIT 10;
	`

	rows, err := database.DB.Query(query, channelName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	commonWords := make(map[string]int)
	for rows.Next() {
		var word string
		var count int
		if err := rows.Scan(&word, &count); err != nil {
			return nil, err
		}
		commonWords[word] = count
	}

	return commonWords, nil
}

// FetchTopVideosByChannel retrieves the top videos from a specific channel.
func FetchTopVideosByChannel(channelName string) ([]string, error) {
	query := `
		SELECT vd.video_id
		FROM video_details vd
		JOIN youtube_channels yc ON vd.channel_id = yc.id
		WHERE yc.channel_name = ?
		ORDER BY (SELECT SUM(count) FROM word_counts WHERE video_id = vd.id) DESC
		LIMIT 10;
	`

	rows, err := database.DB.Query(query, channelName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var topVideos []string
	for rows.Next() {
		var videoID string
		if err := rows.Scan(&videoID); err != nil {
			return nil, err
		}
		topVideos = append(topVideos, videoID)
	}

	return topVideos, nil
}

// FetchVideosByKeyword retrieves videos from a specific channel that contain a particular keyword.
func FetchVideosByKeyword(channelName, keyword string) ([]VideoWordCount, error) {
	query := `
       SELECT vd.video_id, wc.count
       FROM video_details vd
       JOIN word_counts wc ON vd.id = wc.video_id
       JOIN dictionary d ON wc.word_id = d.id
       JOIN youtube_channels yc ON vd.channel_id = yc.id
       WHERE yc.channel_name = ? AND d.word = ?
    `

	rows, err := database.DB.Query(query, channelName, keyword)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var videoWordCounts []VideoWordCount
	for rows.Next() {
		var videoID string
		var count int
		if err := rows.Scan(&videoID, &count); err != nil {
			return nil, err
		}
		videoWordCounts = append(videoWordCounts, VideoWordCount{
			VideoID: videoID,
			Count:   count,
		})
	}

	return videoWordCounts, nil
}

type VideoWordCount struct {
	VideoID string `json:"video_id"`
	Count   int    `json:"count"`
}
type Video struct {
	ID         int         `json:"id"`
	ChannelID  int         `json:"channel_id"`
	VideoID    string      `json:"video_id"`
	WordCounts map[int]int `json:"word_counts"`
}

func InsertVideo(channelID int, videoID string, wordCounts map[int]int) error {
	res, err := database.DB.Exec("INSERT INTO video_details (channel_id, video_id) VALUES (?, ?)", channelID, videoID)
	if err != nil {
		return err
	}
	videoRowID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	for wordID, count := range wordCounts {
		_, err := database.DB.Exec("INSERT INTO word_counts (video_id, word_id, count) VALUES (?, ?, ?)", videoRowID, wordID, count)
		if err != nil {
			return err
		}
	}

	return nil
}

// FetchVideosWithKeywordAcrossChannels retrieves videos that contain a particular keyword across multiple channels.
func FetchVideosWithKeywordAcrossChannels(channelNames []string, keyword string) ([]VideoWordCount, error) {
	query := `
       SELECT vd.video_id, wc.count
       FROM video_details vd
       JOIN word_counts wc ON vd.id = wc.video_id
       JOIN dictionary d ON wc.word_id = d.id
       JOIN youtube_channels yc ON vd.channel_id = yc.id
       WHERE yc.channel_name IN (?)
       AND d.word = ?
    `

	// Convert the slice of channel names to a comma-separated string
	channelNamesStr := "'" + strings.Join(channelNames, "','") + "'"

	rows, err := database.DB.Query(query, channelNamesStr, keyword)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var videoWordCounts []VideoWordCount
	for rows.Next() {
		var videoID string
		var count int
		if err := rows.Scan(&videoID, &count); err != nil {
			return nil, err
		}
		videoWordCounts = append(videoWordCounts, VideoWordCount{
			VideoID: videoID,
			Count:   count,
		})
	}

	return videoWordCounts, nil
}

// FetchVideosWithKeywordAcrossVideos retrieves videos that contain a particular keyword across multiple video IDs.
func FetchVideosWithKeywordAcrossVideos(videoIDs []string, keyword string) ([]VideoWordCount, error) {
	query := `
       SELECT vd.video_id, wc.count
       FROM video_details vd
       JOIN word_counts wc ON vd.id = wc.video_id
       JOIN dictionary d ON wc.word_id = d.id
       WHERE vd.video_id IN (?)
       AND d.word = ?
    `

	// Convert the slice of video IDs to a comma-separated string
	videoIDsStr := "'" + strings.Join(videoIDs, "','") + "'"

	rows, err := database.DB.Query(query, videoIDsStr, keyword)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var videoWordCounts []VideoWordCount
	for rows.Next() {
		var videoID string
		var count int
		if err := rows.Scan(&videoID, &count); err != nil {
			return nil, err
		}
		videoWordCounts = append(videoWordCounts, VideoWordCount{
			VideoID: videoID,
			Count:   count,
		})
	}

	return videoWordCounts, nil
}

// FetchVideosWithKeywordAcrossDB retrieves videos that contain a particular keyword across all channels.
func FetchVideosWithKeywordAcrossDB(keyword string) ([]VideoWordCount, error) {
	query := `
       SELECT vd.video_id, wc.count
       FROM video_details vd
       JOIN word_counts wc ON vd.id = wc.video_id
       JOIN dictionary d ON wc.word_id = d.id
       WHERE d.word = ?
    `

	rows, err := database.DB.Query(query, keyword)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var videoWordCounts []VideoWordCount
	for rows.Next() {
		var videoID string
		var count int
		if err := rows.Scan(&videoID, &count); err != nil {
			return nil, err
		}
		videoWordCounts = append(videoWordCounts, VideoWordCount{
			VideoID: videoID,
			Count:   count,
		})
	}

	return videoWordCounts, nil
}
