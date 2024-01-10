package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Username string `json:"username" binding:"required" gorm:"unique"`
	Email    string `json:"email" binding:"required" gorm:"unique"`
	Password string `json:"password"   binding:"required"`
}
