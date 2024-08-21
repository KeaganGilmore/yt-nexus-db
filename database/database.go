package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("sqlite", "./data/yt-nexus.db?_timeout=10000")
	if err != nil {
		log.Fatal(err)
	}

	createTables()
}
func createTables() {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS dictionary (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            word TEXT UNIQUE
        );`,
		`CREATE TABLE IF NOT EXISTS youtube_channels (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            channel_name TEXT UNIQUE
        );`,
		`CREATE TABLE IF NOT EXISTS video_details (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            channel_id INTEGER,
            video_id TEXT UNIQUE
        );`,
		`CREATE TABLE IF NOT EXISTS word_counts (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            video_id INTEGER,
            word_id INTEGER,
            count INTEGER
        );`,
	}

	for _, query := range queries {
		_, err := DB.Exec(query)
		if err != nil {
			log.Fatalf("Failed to execute query: %v\n", err)
		}
	}
}
