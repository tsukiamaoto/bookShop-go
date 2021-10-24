package model

// user model
type User struct {
	ID       uint   `gorm:primaryKey`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
