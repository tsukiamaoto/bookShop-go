package db

import (
	"fmt"
	"shopCart/config"
	"shopCart/model"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnect(conf *config.Config) *gorm.DB {

	// connected to postgres db just to be able create db statement
	postgresDB, err := gorm.Open(postgres.Open(conf.Databases["default"].Source))
	if err != nil {
		fmt.Println("使用 gorm 連線 DB 發生錯誤，原因為", err)
		log.Error(err)
	}

	// created traget database and connection to target
	dbExec := fmt.Sprintf("CREATE DATABASE %s;", conf.Databases["shopCart"].Name)
	if err := postgresDB.Exec(dbExec); err != nil {
		fmt.Printf("無法建立 %s 資料庫，連線該資料庫\n", conf.Databases["shopCart"].Name)
		conn, err := gorm.Open(postgres.Open(conf.Databases["shopCart"].Source))
		if err != nil {
			fmt.Println("使用 gorm 連線 DB 發生錯誤，原因為", err)
			log.Error(err)
		}

		return conn
	}

	return nil
}

func AutoMigrate(db *gorm.DB) {
	if err := db.AutoMigrate(new(*model.User)); err != nil {
		panic("資料庫User migration的失敗原因是" + err.Error())
	}
	fmt.Println("user db migration 成功！")

	if err := db.AutoMigrate(new(*model.Category)); err != nil {
		panic("Category migration的失敗原因是" + err.Error())
	}
	fmt.Println("category db migration 成功！")

	if err := db.AutoMigrate(new(*model.Product)); err != nil {
		panic("資料庫Product migration的失敗原因是" + err.Error())
	}
	fmt.Println("product db migration 成功！")

	if err := db.AutoMigrate(new(*model.Cart)); err != nil {
		panic("資料庫Cart migration的失敗原因是" + err.Error())
	}
	fmt.Println("cart db migration 成功！")

	if err := db.AutoMigrate(new(*model.CartItem)); err != nil {
		panic("資料庫CartItem migration的失敗原因是" + err.Error())
	}
	fmt.Println("cartItem db migration 成功！")

	if err := db.AutoMigrate(new(*model.Order)); err != nil {
		panic("資料庫Order migration的失敗原因是" + err.Error())
	}
	fmt.Println("order db migration 成功！")

	if err := db.AutoMigrate(new(*model.OrderItem)); err != nil {
		panic("資料庫OrderItem migration的失敗原因是" + err.Error())
	}
	fmt.Println("orderItem db migration 成功！")

	if err := db.AutoMigrate(new(*model.Seller)); err != nil {
		panic("資料庫Seller migration的失敗原因是" + err.Error())
	}
	fmt.Println("seller db migration 成功！")
}
