package repo

import (
	"fmt"
	"github/vamshi1997/projects/go/youtube_app/internal/boot"
	"github/vamshi1997/projects/go/youtube_app/internal/models"

	"gorm.io/gorm"
)

func Create(video models.Video) {
	db := boot.GetDB()

	if err := db.Where("video_id = ?", video.VideoID).First(&video).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			db.Create(&video)
			fmt.Println("Video record created successfully")
		} else {
			fmt.Println("Error querying database:", err)
		}
	} else {
		// Video already exists
		fmt.Println("Video already exists for ID:", video.ID)
	}
}

func FetchVideos(pageSize int, offset int) []models.Video {
	var videos []models.Video

	db := boot.GetDB()

	db.Limit(pageSize).Offset(offset).Order("published_at desc").Find(&videos)

	return videos
}

func SearchVideos(title string, description string) []models.Video {
	var videos []models.Video

	db := boot.GetDB()

	if title != "" {
		db.Where("title LIKE ?", "%"+title+"%").Order("published_at desc").Find(&videos)
		return videos
	}

	if description != "" {
		db.Where("description LIKE ?", "%"+description+"%").Order("published_at desc").Find(&videos)
	}

	return videos
}
