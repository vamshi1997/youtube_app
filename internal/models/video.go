package models

import (
	"time"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	VideoID      string    `json:"video_id" gorm:"index"`
	Title        string    `json:"title" gorm:"index"`
	Description  string    `json:"description" gorm:"index"`
	PublishedAt  time.Time `json:"published_at" gorm:"index"`
	ThumbnailUrl string    `json:"thumbnail_url"`
}
