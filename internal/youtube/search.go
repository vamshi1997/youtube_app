package youtube

import (
	"flag"
	"fmt"
	"github/vamshi1997/projects/go/youtube_app/internal/models"
	"github/vamshi1997/projects/go/youtube_app/internal/repo"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type Request struct {
	Type           string
	Order          string
	PublishedAfter string
	MaxResults     int
	RegionCode     string
	SafeSearch     string
}

const developerKey = "AIzaSyBsYk2-Htpk80WPfUZaI5K61S6yyf6DQQU"

func Search(request Request) {
	flag.Parse()

	service, err := youtube.NewService(&gin.Context{}, option.WithAPIKey(developerKey))
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	// Make the API call to YouTube.
	call := service.Search.List([]string{"id,snippet"}).
		Type(request.Type).
		Order(request.Order).
		RegionCode(request.RegionCode).
		SafeSearch(request.SafeSearch).
		PublishedAfter(request.PublishedAfter).
		MaxResults(int64(request.MaxResults))
	response, err := call.Do()
	if err != nil {
		panic(err)
	}

	// Group video, channel, and playlist results in separate lists.
	videos := []models.Video{}

	// Iterate through each item and add it to the correct list.
	for _, item := range response.Items {

		PublishedAt, err := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
		if err != nil {
			fmt.Println("Error:", err)
		}

		VideoModel := models.Video{
			VideoID:      item.Id.VideoId,
			Title:        item.Snippet.Title,
			Description:  item.Snippet.Description,
			PublishedAt:  PublishedAt,
			ThumbnailUrl: item.Snippet.Thumbnails.Default.Url,
		}

		videos = append(videos, VideoModel)
	}

	for _, video := range videos {
		repo.Create(video)
	}

	fmt.Println("updated latest videos in DB...")
}
