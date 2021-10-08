package model

// user model
type User struct {
	ID       uint   `gorm:primaryKey`
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
