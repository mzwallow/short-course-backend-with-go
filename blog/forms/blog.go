// Package forms contains the request and response structs for the blog API.
// The request and response structs are used to validate the request and
// format the response.
// The request and response structs are used to control how the data is
// serialized to JSON.
package forms

import "time"

// CreateBlogRequest represents a request to create a blog.
//
// It contains the title and content of the blog.
//
// The binding:"required" tag is used to validate the request. If the title or
// content fields are empty, the request will be rejected with a 400 Bad Request
// response.
//
// For more information on request validation, see:
// https://echo.labstack.com/guide/request
type CreateBlogRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// GetAllBlogsResponse represents a response containing a list of blogs.
//
// It contains the ID, title, content, created_at, updated_at and comments
// fields.
//
// The comments field is the number of comments that the blog has.
//
// It has json struct tags, which are used to control how the fields are
// serialized to JSON.
//
// For more information on JSON struct tags, see:
// https://golang.org/pkg/encoding/json/#Marshal
type GetAllBlogsResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Comments  int       `json:"comments"`
}

// GetBlogByIDResponse represents a response containing a single blog.
//
// It embeds the Blog struct, which means that it has all the same fields as
// the Blog struct.
//
// It also has a Comments field, which is a slice of Comment structs.
//
// This is an example of composition in Go.
//
// For more information on composition in Go, see:
// https://golang.org/doc/effective_go.html#embedding
type GetBlogByIDResponse struct {
	Blog
	Comments []Comment `json:"comments"`
}

// UpdateBlogRequest represents a request to update a blog.
//
// It contains the title and content of the blog.
type UpdateBlogRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// CreateCommentRequest represents a request to create a comment.
//
// It contains the blog_id and content of the comment.
type CreateCommentRequest struct {
	Content string `json:"content" binding:"required"`
}

// Blog represents a blog in the database.
type Blog struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Comment represents a comment in the database.
type Comment struct {
	ID        int       `json:"id"`
	BlogID    int       `json:"blog_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
