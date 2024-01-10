package models

import (
	"blogApp/api/database"

	"golang.org/x/crypto/bcrypt"
)

// ilişkisellik örnek
type Bir struct {
	Bid uint `gorm:"primaryKey"`
	Cok Cok  `gorm:"foreignKey:Bfk"`
}
type Cok struct {
	Cid uint `gorm:"primaryKey"`
	Bfk uint
}

// user
func (user *User) CreateUserRecord() error {
	result := database.GlobalDB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Blog
// create
func (blog *Blog) CreateBlogRecord() error {
	result := database.GlobalDB.Create(&blog)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// delete comment
func (deleteComment *Comment) DeleteComment(id int) error {
	db := database.GlobalDB

	err := db.Unscoped().Where("comid =? ", id).Delete(&deleteComment)
	if err != nil {
		return err.Error
	}
	return nil
}

// delete
func (deleteblog *Blog) DeletePost(id int) error {
	db := database.GlobalDB

	err := db.Unscoped().Where("pid =? ", id).Delete(&deleteblog)
	if err != nil {
		return err.Error
	}
	return nil
}

// update

func (now *Blog) UpdateBlog(updateted *Blog, id int) error {

	db := database.GlobalDB
	now.Header = updateted.Header
	now.Code1 = updateted.Code1
	now.Content1 = updateted.Content1
	now.Userid = updateted.Userid
	err := db.Where("pid= ?", id).First(&now)
	if now.Pid == 0 {
		return err.Error
	}
	err = db.Where("pid= ?", id).Updates(&now)
	if err != nil {
		return err.Error
	}
	return nil

}

func (catpost *Catepostrel) CreateCatPost() error {
	result := database.GlobalDB.Create(&catpost)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// comment
func (comment *Comment) CreateCommentRecord() error {
	result := database.GlobalDB.Create(&comment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// emote
func (emote *Emote) CreateEmoteRecord() error {
	result := database.GlobalDB.Create(&emote)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// hash
func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
