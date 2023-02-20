// Package controllers contains the controllers for the blog API.
// The controllers are responsible for handling requests and returning responses.
// The controllers are the middle layer between the models and the server.
// The controllers are responsible for validating the request data and calling the appropriate model functions.
// The controllers are responsible for formatting the response data.
// The controllers are responsible for returning the appropriate HTTP status code.
package controllers

import (
	"net/http"
	"strconv"

	"blog/forms"
	"blog/models"

	"github.com/gin-gonic/gin"
)

// BlogController is a controller for the blog resource.
type BlogController struct {
	blogModel *models.BlogModel
}

// NewBlogController creates a new BlogController.
func NewBlogController(blogModel *models.BlogModel) *BlogController {
	return &BlogController{
		blogModel: blogModel,
	}
}

// Health returns a 200 OK response if the service is healthy.
func (c *BlogController) Health(ctx *gin.Context) {
	// ctx.JSON is a helper function provided by Gin to write JSON responses.
	// It takes the HTTP status code and a data object as arguments.
	//
	// For more information on ctx.JSON, see:
	// https://godoc.org/github.com/gin-gonic/gin#Context.JSON
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

// CreateBlog creates a new blog.
func (c *BlogController) CreateBlog(ctx *gin.Context) {
	// Create a new instance of the CreateBlogRequest struct.
	var req forms.CreateBlogRequest
	// ctx.BindJSON is a helper function provided by Gin to bind the request body
	// to a Go struct.
	//
	// It takes a pointer to a struct as an argument. The struct must be a
	// pointer, otherwise the request will fail with a 500 Internal Server Error
	// response.
	//
	// For more information on ctx.BindJSON, see:
	// https://godoc.org/github.com/gin-gonic/gin#Context.BindJSON
	if err := ctx.BindJSON(&req); err != nil {
		// If the request fails, we return a 400 Bad Request response.
		//
		// For more information on HTTP status codes, see:
		// https://en.wikipedia.org/wiki/List_of_HTTP_status_codes
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create blog: " + err.Error(),
		})
		return
	}

	// Call the CreateBlog method on the BlogModel, passing in the request data.
	// If the method returns an error, we return a 500 Internal Server Error
	// response.
	//
	// For more information on c.blogModel.CreateBlog, see:
	// blog/models/blog.go
	if err := c.blogModel.CreateBlog(req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create blog: " + err.Error(),
		})
		return
	}

	// If the method succeeds, we return a 200 OK response.
	ctx.JSON(http.StatusOK, gin.H{"message": "Blog created successfully"})
}

// GetAllBlogs returns a list of all blogs.
func (c *BlogController) GetAllBlogs(ctx *gin.Context) {
	// Call the GetAllBlogs method on the BlogModel.
	blogs, err := c.blogModel.GetAllBlogs()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get all blogs: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, blogs)
}

// GetBlogByID returns a single blog.
func (c *BlogController) GetBlogByID(ctx *gin.Context) {
	// ctx.Param is a helper function provided by Gin to get a URL parameter.
	//
	// For more information on ctx.Param, see:
	// https://godoc.org/github.com/gin-gonic/gin#Context.Param
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get blog: " + err.Error(),
		})
		return
	}

	// Call the GetBlogByID method on the BlogModel, passing in the ID.
	blog, err := c.blogModel.GetBlogByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get blog: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, blog)
}

// UpdateBlog updates a blog.
func (c *BlogController) UpdateBlog(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update blog: " + err.Error(),
		})
		return
	}

	var req forms.UpdateBlogRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update blog: " + err.Error(),
		})
		return
	}

	// Call the UpdateBlog method on the BlogModel, passing in the ID and
	// request data.
	if err := c.blogModel.UpdateBlog(id, req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update blog: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Blog updated successfully"})
}

// DeleteBlog deletes a blog.
func (c *BlogController) DeleteBlog(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to delete blog: " + err.Error(),
		})
		return
	}

	// Call the DeleteBlog method on the BlogModel, passing in the ID.
	if err := c.blogModel.DeleteBlog(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete blog: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
}

// CreateComment creates a new comment.
func (c *BlogController) CreateComment(ctx *gin.Context) {
	blogIDString := ctx.Param("id")
	blogID, err := strconv.Atoi(blogIDString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create comment: " + err.Error(),
		})
		return
	}

	var req forms.CreateCommentRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create comment: " + err.Error(),
		})
		return
	}

	// Call the CreateComment method on the BlogModel, passing in the blog ID
	// and request data.
	if err := c.blogModel.CreateComment(blogID, req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create comment: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Comment created successfully"})
}
