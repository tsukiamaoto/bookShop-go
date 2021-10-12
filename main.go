package main

import (
	"test/db"
	"test/module/user/delivery/http"
	"test/module/user/repository"
	"test/module/user/service"
	"test/util"

	"github.com/gin-gonic/gin"
)

func main() {
	// load config
	config := util.LoadConfig()

	// create server
	server := gin.Default()
	// get html page
	server.LoadHTMLGlob("template/html/*")
	// get static resources
	server.Static("/assets", "./template/assets")

	// connect database
	conn := db.DbConnect(config.DBSource)
	db.AutoMigrate(conn)

	// create repository instance
	userRepo := repository.NewUserRepository(conn)

	// create service instance
	userService := service.NewUserSerivce(userRepo)

	// create delivery instance
	http.NewUserHttpHandler(server, userService)

	// start server
	server.Run(config.ServerAddress)
}
