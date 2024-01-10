package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	Comid     uint   `gorm:"primaryKey"`
	Email     string `gorm:"not null" binding:"required"`
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *gorm.DeletedAt `gorm:"index"`
	Blogid    uint
	Blog      Blog `gorm:"foreignKey:Blogid"`
	Eidfk     *uint
	Emote     Emote `gorm:"foreignKey:Eidfk"`
}
