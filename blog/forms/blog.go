package forms

import (
	"time"
)

type Blog struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Comment struct {
	ID        int    `json:"id"`
	BlogID    int    `json:"blog_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateBlogRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type UpdateBlogRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type GetAllBlogsResponse struct {
	Blog
	Comments int `json:"comments"`
}

type GetBlogResponse struct {
	Blog
	Comments []*Comment `json:"comments"`
}

type CreateCommentRequest struct {
	Content string `json:"content" binding:"required"`
}
