package main

import (
	"net/http"
	"errors"
	"github.com/gin-gonic/gin"
)

// page handler
func loginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func loginAuth(c *gin.Context) {
	var (
		username string
		password string
	)

	if in, isExist := c.GetPostForm("username"); isExist && in != "" {
		username = in
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": errors.New("必須輸入使用者名稱"),
		})
		return
	}

	if in, isExist := c.GetPostForm("password"); isExist && in != "" {
		password = in
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": errors.New("必須輸入密碼"),
		})
		return
	}

	if err := Auth(username, password); err == nil {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"success": "成功登入",
		})
		return
	} else {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": err,
		})
		return
	}
}

func main() {
	// create server
	server := gin.Default()
	// get html page
	server.LoadHTMLGlob("template/html/*")
	// get static resources
	server.Static("/assets", "./template/assets")

	// router setting
	server.GET("/login", loginPage)
	server.POST("/login", loginAuth)

	// start server
	server.Run(":9999")
}
