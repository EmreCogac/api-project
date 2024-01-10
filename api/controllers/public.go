package controllers

import (
	"blogApp/api/auth"
	"blogApp/api/database"
	"blogApp/api/models"
	"log"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginPayload struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshtoken"`
}

type ParamID struct {
	Id int `json:"id" binding:"required"`
}

func Signup(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"Error": "Invalid Inputs ",
		})
		c.Abort()
		return
	}
	err = user.HashPassword(user.Password)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"Error": "Error Hashing Password",
		})
		c.Abort()
		return
	}

	err = user.CreateUserRecord()
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"Error": "Error Creating User",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"Message": "Sucessfully Register",
	})
}

func Login(c *gin.Context) {
	var payload LoginPayload
	var user models.User
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Invalid Inputs",
		})
		c.Abort()
		return
	}
	result := database.GlobalDB.Where("email = ?", payload.Email).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(401, gin.H{
			"Error": "Invalid User Credentials",
		})
		c.Abort()
		return
	}

	err = user.CheckPassword(payload.Password)
	if err != nil {
		log.Println(err)
		c.JSON(401, gin.H{
			"Error": "Invalid User Credentials",
		})
		c.Abort()
		return
	}
	jwtWrapper := auth.JwtWrapper{
		SecretKey:         "verysecretkey",
		Issuer:            "AuthService",
		ExpirationMinutes: 1,
		ExpirationHours:   12,
	}
	signedToken, err := jwtWrapper.GenerateToken(user.Email)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"Error": "Error Signing Token",
		})
		c.Abort()
		return
	}
	signedtoken, err := jwtWrapper.RefreshToken(user.Email)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"Error": "Error Signing Token",
		})
		c.Abort()
		return
	}
	tokenResponse := LoginResponse{
		Token:        signedToken,
		RefreshToken: signedtoken,
	}
	c.JSON(200, tokenResponse)
}

// blogs
// update
func UpdatePost(c *gin.Context) {
	var updatePost models.Blog
	var blog models.Blog

	err := c.ShouldBindJSON(&updatePost)
	if err != nil {
		c.JSON(400, gin.H{
			"err": "invalid ",
		})
		c.Abort()
	}
	err = blog.UpdateBlog(&updatePost, int(updatePost.Pid))
	if err != nil {
		c.JSON(400, gin.H{
			"err": "update problem ",
		})
		c.Abort()
	} else {
		c.JSON(200, gin.H{
			"işlem başarılı": "işlem başarılı",
		})
	}

}

// delete
func DeleteBlog(c *gin.Context) {
	var ID ParamID
	var blog models.Blog
	db := database.GlobalDB
	err := c.ShouldBindJSON(&ID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid id ",
		})

		c.Abort()
		return
	}
	db.Where("pid =? ", ID.Id).Find(&blog)

	if blog.Pid == 0 {
		log.Println("cant resolve id")
		c.JSON(500, gin.H{
			" err ": " cant find id ",
			"id":    ID.Id,
		})
		c.Abort()
		return

	}
	err = blog.DeletePost(ID.Id)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			" err ": " error deleting post ",
		})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"sucsess": " deleted ",
	})

}

func DeleteComment(c *gin.Context) {
	var ID ParamID
	var comment models.Comment
	db := database.GlobalDB
	err := c.ShouldBindJSON(&ID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid id ",
		})

		c.Abort()
		return
	}
	db.Where("comid =? ", ID.Id).Find(&comment)

	if comment.Comid == 0 {
		log.Println("cant resolve id")
		c.JSON(500, gin.H{
			" err ": " cant find id ",
			"id":    ID.Id,
		})
		c.Abort()
		return

	}
	err = comment.DeleteComment(ID.Id)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			" err ": " error deleting post ",
		})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"sucsess": " deleted ",
	})

}

// blog create
func CreatePost(c *gin.Context) {
	var post models.Blog
	err := c.ShouldBindJSON(&post)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"err": "eksik girdi",
		})

		c.Abort()
		return
	}
	err = post.CreateBlogRecord()
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"err": "Error post oluşturulamadi	",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"Message": "oluşturuldu",
	})

}

func Goster(c *gin.Context) {
	db := database.GlobalDB
	var cate []models.Catepostrel

	db.Preload("Blog").Find(&cate)

	c.JSON(200, gin.H{
		"data": cate,
	})

}

func CreateCatPost(c *gin.Context) {
	var relCat models.Catepostrel
	err := c.ShouldBindJSON(&relCat)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"err": "eksik girdi",
		})

		c.Abort()
		return
	}
	err = relCat.CreateCatPost()
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"err": "Error post oluşturulamadi",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"Message": "oluşturuldu",
	})

}

// get posts
func GetAllBlogs(c *gin.Context) {
	pageStr := c.Query("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, _ := strconv.Atoi(pageStr)
	limit := 5
	offsets := (page - 1) * limit
	var total int64

	// var getBlog []models.Blog
	var blogpost []models.Blog
	db := database.GlobalDB
	// db.Preload("Blog").Preload("Categories").Preload("Blog.User").Offset(offsets).Limit(limit).Find(&catpost)
	// db.Model(&models.Blog{}).Count(&total)
	db.Preload("Blog.User").Offset(offsets).Limit(limit).Find(&blogpost)
	db.Model(&models.Blog{}).Count(&total)

	c.JSON(200, gin.H{
		"data": blogpost,
		"meta": gin.H{
			"total":     total,
			"page ":     page,
			"last_page": math.Ceil(float64(total) / float64(limit)),
		},
	})

}

func CommentPost(c *gin.Context) {
	db := database.GlobalDB
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {

		c.JSON(400, gin.H{"error": "Invalid ID parameter"})
		return
	}
	var comment models.Comment
	db.Where("blogid =?", id).Preload("Blog").Find(&comment)
	c.JSON(200, gin.H{
		"data":  comment,
		"param": id,
	})
}

// post info
func PostDetail(c *gin.Context) {
	db := database.GlobalDB
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {

		c.JSON(400, gin.H{"error": "Invalid ID parameter"})
		return
	}
	var blogpost models.Blog
	db.Where("pid =?", id).Preload("User").Find(&blogpost)
	c.JSON(200, gin.H{
		"data":  blogpost,
		"param": id,
	})
}

// comment create
func CreateComment(c *gin.Context) {
	var comment models.Comment
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"err": "eksik girdi",
		})

		c.Abort()
		return
	}
	err = comment.CreateCommentRecord()
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"err": "Error post oluşturulamadi	",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"Message": "oluşturuldu",
	})

}

// func GetAll(c *gin.Context) {
// 	db := database.GlobalDB

// 	var posts []models.Posts

// 	db.Preload("User").Find(&posts, models.Posts.User.)

// 	c.JSON(200, gin.H{
// 		"posts": posts,
// 	})

// }
