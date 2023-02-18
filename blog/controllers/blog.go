package controllers

import (
	"net/http"
	"strconv"

	"blog/forms"
	"blog/models"

	"github.com/gin-gonic/gin"
)

type BlogControllers struct {
	blogModels *models.BlogModels
}

func NewBlogControllers(blogModels *models.BlogModels) *BlogControllers {
	return &BlogControllers{blogModels: blogModels}
}

func (c *BlogControllers) CreateBlog(ctx *gin.Context) {
	var req forms.CreateBlogRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	blog := &forms.Blog{
		Title:   req.Title,
		Content: req.Content,
	}

	if err := c.blogModels.CreateBlog(blog); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Blog created successfully"})
}

func (c *BlogControllers) GetAllBlogs(ctx *gin.Context) {
	posts, err := c.blogModels.GetAllBlogs()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, posts)
}

func (c *BlogControllers) GetBlog(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	blog, err := c.blogModels.GetBlog(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, blog)
}

func (c *BlogControllers) UpdateBlog(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req forms.UpdateBlogRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	blog := &forms.Blog{
		ID:      id,
		Title:   req.Title,
		Content: req.Content,
	}

	if err := c.blogModels.UpdateBlog(blog); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Blog updated successfully"})
}

func (c *BlogControllers) DeleteBlog(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.blogModels.DeleteBlog(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
}

func (c *BlogControllers) CreateComment(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req forms.CreateCommentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment := &forms.Comment{
		BlogID:  id,
		Content: req.Content,
	}

	if err := c.blogModels.CreateComment(comment); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Comment created successfully"})
}

func (c *BlogControllers) DeleteComment(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.blogModels.DeleteComment(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
