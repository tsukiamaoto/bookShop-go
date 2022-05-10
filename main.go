package main

import (
	"github.com/tsukiamaoto/bookShop-go/config"
	"github.com/tsukiamaoto/bookShop-go/db"
	"github.com/tsukiamaoto/bookShop-go/module/delivery"
	"github.com/tsukiamaoto/bookShop-go/module/repository"
	"github.com/tsukiamaoto/bookShop-go/module/server"
	"github.com/tsukiamaoto/bookShop-go/module/service"
	"github.com/tsukiamaoto/bookShop-go/redis"
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
	services := service.NewServices(service.Repos{
		Repos: repos,
	})

	// create delivery instance
	handler := delivery.NewHandler(services)

	// create server isntance
	srv := server.NewServer(config, handler.Init(config))
	// start server
	srv.Run()
}
