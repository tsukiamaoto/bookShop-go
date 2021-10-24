package main

import (
	"shopCart/config"
	"shopCart/db"
	"shopCart/module/user/delivery/http"
	"shopCart/module/user/repository"
	"shopCart/module/user/service"
	"shopCart/redis"
	
	"github.com/gin-gonic/gin"
)

func init() {
	redis.ConnectRDB()
}

func main() {
	// load config
	config := config.LoadConfig()

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
