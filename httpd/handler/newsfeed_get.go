package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"learn/platform/newsfeed"
)

func NewsFeedGet(f *newsfeed.Repo) gin.HandlerFunc{
	return func(c *gin.Context){
		 c.JSON(http.StatusOK, f.GetAll())
	}
}