package db

import (
	"time"

	"gorm.io/gorm"
)

type Items struct {
	gorm.Model
	ItemID   int `gorm:"primarykey"`
	Title    string
	Album    string
	Artist   string
	Url      string
	AcrID    string
	Provider string
	Metadata ItemMetadata `gorm:"embedded"`
	Comment  Comments     `gorm:"foreignKey:ItemID;references:ItemID"`
}

type ItemMetadata struct {
	SpotifyURL string
	SpotifyID  string
	YoutubeURL string
	YoutubeID  string
}

type Comments struct {
	gorm.Model
	CommentID int `gorm:"unique;primaryKey"`
	Up        int
	Down      int
	Content   string
	Created   *time.Time
	ItemID    int `gorm:"not null"`
	Thumb     string
}
