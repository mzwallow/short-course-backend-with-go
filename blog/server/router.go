package server

import (
	"blog/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(blogCtrls *controllers.BlogControllers) *gin.Engine {
	r := gin.Default()
	c := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Accept", "X-Requested-With", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}
	r.Use(cors.New(c))

	r.GET("/health", blogCtrls.Health)

	blogs := r.Group("/blogs")
	{
		blogs.POST("/", blogCtrls.CreateBlog)
		blogs.GET("/", blogCtrls.GetAllBlogs)
		blogs.GET("/:id", blogCtrls.GetBlog)
		blogs.PUT("/:id", blogCtrls.UpdateBlog)
		blogs.DELETE("/:id", blogCtrls.DeleteBlog)
		blogs.POST("/:id/comments", blogCtrls.CreateComment)
		blogs.DELETE("/:id/comments/:comment_id", blogCtrls.DeleteComment)
	}

	return r
}
