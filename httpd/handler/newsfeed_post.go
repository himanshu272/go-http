package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"learn/platform/newsfeed"
)

type request struct {
	Title string `json:"title"`
	Post string `json:"post"`
	URL string `json:"url"`
}

func NewsFeedPost(f *newsfeed.Repo) gin.HandlerFunc{
	return func(c *gin.Context){
		requestBody := request{}
		c.Bind(&requestBody)

		item := newsfeed.Item{
			Title: requestBody.Title,
			Post: requestBody.Post,
		}

		f.Add(item)

		c.Status(http.StatusNoContent)
	}
}