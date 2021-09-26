package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// init database
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = 9955
	dbname   = "demo"
)

// user model
type User struct {
	ID       uint `gorm:primaryKey`
	Username string
	Password string
}

func dbConnect() *gorm.DB {
	// connect db
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%d dbname=%s sslmode=disable", host, port, user, password, dbname)
	// open db
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("使用 gorm 連線 DB 發生錯誤，原因為" + err.Error())
	}

	return db
}

func autoMigrate(db *gorm.DB) {
	if err := db.AutoMigrate(new(User)); err != nil {
		panic("資料庫migration的失敗原因是" + err.Error())
	}
	fmt.Println(" migration 成功！")
}

func insertUser(db *gorm.DB, user *User) {
	result := db.Create(&user)
	if result.Error != nil {
		panic("建立使用者失敗，原因是" + result.Error.Error())
	}

	fmt.Println("建立使用者成功")
}

func queryUser(db *gorm.DB, username string) {
	var user User
	result := db.Where("Username=?", username).First(&user)

	if result.Error != nil {
		panic("查詢不到結果的原因:" + result.Error.Error())
	}

	fmt.Println("查詢使用者成功:", user)
}

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
	// connect db
	db := dbConnect()
	user := &User{
		Username: "test",
		Password: "test",
	}
	autoMigrate(db)
	insertUser(db, user)
	queryUser(db, "test")

	setConfig()
	/*// create server
	server := gin.Default()
	// get html page
	server.LoadHTMLGlob("template/html/*")
	// get static resources
	server.Static("/assets", "./template/assets")

	// router setting
	server.GET("/login", loginPage)
	server.POST("/login", loginAuth)

	// start server
	server.Run(":9999")*/
}
