package controllers

import (
	"fmt"
	"github/vamshi1997/projects/go/youtube_app/internal/repo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Status(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})
}

func FetchVideos(ctx *gin.Context) {

	page := ctx.DefaultQuery("page", "1")
	pageSize := ctx.DefaultQuery("pageSize", "10")

	pageInt, _ := strconv.Atoi(page)
	pageSizeInt, _ := strconv.Atoi(pageSize)

	offset := (pageInt - 1) * pageSizeInt

	ctx.JSON(http.StatusOK, repo.FetchVideos(pageSizeInt, offset))
}

func SearchVideos(ctx *gin.Context) {

	title := ctx.Query("title")
	description := ctx.Query("description")

	fmt.Print(title)

	ctx.JSON(http.StatusOK, repo.SearchVideos(title, description))
}
