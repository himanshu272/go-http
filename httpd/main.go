package main

import (
	"learn/httpd/handler"
	"github.com/gin-gonic/gin"
	"learn/platform/newsfeed"
)

func main() {
	feed := newsfeed.New()

	r := gin.Default()

	r.GET("/newsfeed", handler.NewsFeedGet(feed))
	r.POST("/newsfeed", handler.NewsFeedPost(feed))
	r.POST("/download", handler.DownloadFile())

	r.Run()
}
