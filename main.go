package main

import (
	"shopCart/config"
	"shopCart/db"
	"shopCart/module/delivery"
	"shopCart/module/repository"
	"shopCart/module/service"
	"shopCart/module/server"
	"shopCart/redis"
)

func init() {
	redis.ConnectRDB()
}

func main() {
	// load config
	config := config.LoadConfig()

	// connect database
	conn := db.DbConnect(config)
	db.AutoMigrate(conn)

	// create repository instance
	repos := repository.NewRepositories(conn)

	// create service instance
	services := service.NewServices(service.Deps{
		Repos: repos,
	})

	// create delivery instance
	handler := delivery.NewHandler(services)
	
	// create server isntance
	srv := server.NewServer(config, handler.Init(config))
	// start server
	srv.Run()
}
