package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// init database
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = 9955
	dbname   = "demo"
)

type User struct {
	ID       string
	Username string
	Password string
}

func dbConnect() *sql.DB {
	// connect db
	connect := fmt.Sprintf("host=%s port=%d user=%s password=%d dbname=%s sslmode=disable", host, port, user, password, dbname)
	// open db
	db, err := sql.Open("postgres", connect)

	if err != nil {
		fmt.Println("開啟 MySQL 連線發生錯誤，原因為：", err)
	} else if err := db.Ping(); err != nil {
		fmt.Println("資料庫連線錯誤，原因為:", err.Error())
	}

	return db
}

func CreateTable(db *sql.DB) error {
	sql := `CREATE TABLE IF NOT EXISTS users(
		id serial NOT NULL PRIMARY KEY,
		username VARCHAR(64),
		password VARCHAR(64)
	); `

	if _, err := db.Exec(sql); err != nil {
		fmt.Println("建立 Table 發生錯誤:", err)
		return err
	}
	fmt.Println("建立 Table 成功！")
	return nil
}

func insertUser(DB *sql.DB, username, password string) error {
	_, err := DB.Exec("insert INTO users(username,password) values($1,$2)", username, password)
	if err != nil {
		fmt.Printf("建立使用者失敗，原因是：%v", err)
		return err
	}

	fmt.Println("建立使用者成功")
	return nil
}

func queryUser(db *sql.DB, username string) {
	user := new(User)
	row := db.QueryRow("select * from users where username=$1", username)

	if err := row.Scan(&user.ID, &user.Username, &user.Password); err != nil {
		fmt.Printf("映射使用者失敗，失敗的原因是:%v\n", err)
		return
	}

	fmt.Println("查詢使用者成功", *user)
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
	CreateTable(db)
	insertUser(db, "test", "test")
	queryUser(db, "test")

	defer db.Close()
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
