// This file contains the router for the blog service.
// It defines the routes for the service and maps them to the appropriate
// controller functions.
package server

import (
	"blog/controllers"

	"github.com/gin-gonic/gin"
)

// NewRouter creates a new router.
func NewRouter(blogCtrl *controllers.BlogController) *gin.Engine {
	// Create a new router.
	r := gin.Default()

	// Register the health route.
	r.GET("/health", blogCtrl.Health)

	// r.GET("/blogs", blogCtrl.GetAllBlogs)
	// r.POST("/blogs", blogCtrl.CreateBlog)
	// r.PUT("/blogs/:id", blogCtrl.UpdateBlog)
	// r.DELETE("/blogs/:id", blogCtrl.DeleteBlog)
	// r.POST("/blogs/:id/comments", blogCtrl.CreateComment)

	// Above is the original code. We can group the routes together to make it more readable.
	// The code below is the same as the code above.
	// For more information on route grouping, see:
	// https://godoc.org/github.com/gin-gonic/gin#RouterGroup
	blogs := r.Group("/blogs")
	{
		blogs.POST("", blogCtrl.CreateBlog)
		blogs.GET("", blogCtrl.GetAllBlogs)
		blogs.GET("/:id", blogCtrl.GetBlogByID)
		blogs.PUT("/:id", blogCtrl.UpdateBlog)
		blogs.DELETE("/:id", blogCtrl.DeleteBlog)
		blogs.POST("/:id/comments", blogCtrl.CreateComment)
	}

	return r
}
