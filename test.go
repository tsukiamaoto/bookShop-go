package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type IndexData struct {
	Title   string
	Content string
}

// html page
func test(c *gin.Context) {
	data := new(IndexData)
	data.Title = "首頁"
	data.Content = "我的第一個首頁"
	c.HTML(http.StatusOK, "index.html", data)
}

func main() {
	// create server
	server := gin.Default()
	server.LoadHTMLGlob("template/*")

	// router setting
	server.GET("/", test)

	// start server
	server.Run(":9999")
}
