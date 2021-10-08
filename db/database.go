package db

import (
	"fmt"
	"test/model"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnect(dbSource string) *gorm.DB {
	// connect db
	dsn := dbSource
	// open db
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("使用 gorm 連線 DB 發生錯誤，原因為" + err.Error())
	}

	return conn
}

func AutoMigrate(db *gorm.DB) {
	if err := db.AutoMigrate(new(*model.User)); err != nil {
		panic("資料庫migration的失敗原因是" + err.Error())
	}
	fmt.Println(" migration 成功！")
}
