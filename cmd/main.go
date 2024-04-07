package main

import (
	"fmt"
	"github/vamshi1997/projects/go/youtube_app/internal/boot"
	"github/vamshi1997/projects/go/youtube_app/internal/router"
	"github/vamshi1997/projects/go/youtube_app/internal/youtube"
	"time"
)

func main() {
	boot.InitApp()
	go fetchDataFromExternalEndpoint()
	router.InitiateRouter()
}

func fetchDataFromExternalEndpoint() {

	for {
		time.Sleep(10 * time.Second)

		fmt.Println("----------------------------")
		fmt.Println("searching latest videos ...")

		dateTime := time.Now().Add(-1 * time.Hour).UTC()
		formattedTime := dateTime.Format(time.RFC3339)

		request := youtube.Request{
			Type:           "video",
			Order:          "date",
			PublishedAfter: formattedTime,
			MaxResults:     100,
			RegionCode:     "IN",
			SafeSearch:     "strict",
		}

		youtube.Search(request)
		fmt.Println("----------------------------")
	}
}
