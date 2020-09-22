package handler

import (
	"path"
	"net/http"
	"os"
	"io"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func DownloadFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := request{}
		c.Bind(&requestBody)

		filePath := os.Getenv("FILEPATH")

		filenameArr := strings.Split(requestBody.URL, "/")
		fileName := filenameArr[len(filenameArr)-1]

		filePath = path.Join(filePath,fileName)

		out, err := os.Create(filePath)
		if err != nil {
			fmt.Errorf("got an error %q", err)
		}
		defer out.Close()

		resp, err := http.Get(requestBody.URL)
		if err != nil {
			fmt.Errorf("got an error %q", err)
		}

		defer resp.Body.Close()

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			fmt.Errorf("got an error %q", err)
		}

		fmt.Println("downloaded!")
	}
}