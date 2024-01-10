package route

import (
	"blogApp/api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/get", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "deneme",
		})
	})

	api := r.Group("/api")
	{
		api.GET("/getcate", controllers.Goster)

		getAll := api.Group("/getAll")
		{
			getAll.GET("/post/", controllers.GetAllBlogs)
			getAll.GET("/post/:id", controllers.PostDetail)
			getAll.GET("/comment/:id", controllers.CommentPost)

		}
		public := api.Group("/public")
		{
			public.POST("/login", controllers.Login)
			public.POST("/signup", controllers.Signup)
			public.POST("/create/comment", controllers.CreateComment)
			public.POST("/delete/comment", controllers.DeleteComment)

		}

		protected := api.Group("/protected") //.Use(middlewares.Authz())
		{
			protected.POST("/addcat/post", controllers.CreateCatPost)
			protected.POST("/deletecat/post")

			protected.POST("/create/post", controllers.CreatePost)
			protected.POST("/update/post", controllers.UpdatePost)
			protected.POST("/delete/post", controllers.DeleteBlog)
			protected.GET("/profile", controllers.Profile)

		}
	}

	return r
}
