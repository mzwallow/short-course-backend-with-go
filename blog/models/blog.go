package models

import (
	"database/sql"
	"fmt"
	"time"

	"blog/forms"
)

type BlogModels struct {
	db *sql.DB
}

func NewBlogModels(db *sql.DB) *BlogModels {
	return &BlogModels{db: db}
}

func (m *BlogModels) CreateBlog(blog *forms.Blog) error {
	// Create a transaction
	tx, err := m.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	// Prepare the statement
	stmt, err := tx.Prepare(`INSERT INTO blogs (title, content) VALUES ($1, $2)`)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %v", err)
	}
	defer stmt.Close()

	// Execute the statement
	if _, err = stmt.Exec(blog.Title, blog.Content); err != nil {
		return fmt.Errorf("failed to execute statement: %v", err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

func (m *BlogModels) GetAllBlogs() ([]*forms.GetAllBlogsResponse, error) {
	rows, err := m.db.Query(
		`SELECT 
			b.id, 
			b.title, 
			SUBSTRING(b.content FROM 0 FOR 350),
			b.created_at, 
			b.updated_at,
			COUNT(c.id) AS comments
		FROM blogs AS b
		LEFT JOIN comments AS c ON b.id = c.blog_id
		GROUP BY b.id
		ORDER BY b.created_at DESC`,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get all blogs: %v", err)
	}
	defer rows.Close()

	var blogs []*forms.GetAllBlogsResponse
	for rows.Next() {
		var blog forms.GetAllBlogsResponse
		if err := rows.Scan(
			&blog.ID,
			&blog.Title,
			&blog.Content,
			&blog.CreatedAt,
			&blog.UpdatedAt,
			&blog.Comments,
		); err != nil {
			return nil, fmt.Errorf("failed to scan blog: %v", err)
		}

		blogs = append(blogs, &blog)
	}

	return blogs, nil
}

func (m *BlogModels) GetBlog(id int) (*forms.GetBlogResponse, error) {
	blog := &forms.GetBlogResponse{}
	if err := m.db.QueryRow(
		`SELECT id,
				title,
				content,
				created_at,
				updated_at
		FROM blogs WHERE id = $1`,
		id,
	).Scan(
		&blog.ID,
		&blog.Title,
		&blog.Content,
		&blog.CreatedAt,
		&blog.UpdatedAt,
	); err != nil {
		return nil, fmt.Errorf("failed to get blog: %v", err)
	}

	// Get the comments for the blog
	rows, err := m.db.Query(
		`SELECT id,
				blog_id,
				content,
				created_at,
				updated_at
		FROM comments WHERE blog_id = $1`,
		id,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get comments for blog: %v", err)
	}
	defer rows.Close()

	var comments []*forms.Comment
	for rows.Next() {
		var comment forms.Comment
		if err := rows.Scan(
			&comment.ID,
			&comment.BlogID,
			&comment.Content,
			&comment.CreatedAt,
			&comment.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan comment: %v", err)
		}

		comments = append(comments, &comment)
	}

	blog.Comments = comments

	return blog, nil
}

func (m *BlogModels) UpdateBlog(blog *forms.Blog) error {
	// Create a transaction
	tx, err := m.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	// Prepare the statement
	stmt, err := tx.Prepare(`UPDATE blogs SET title = $1, content = $2, updated_at = $3 WHERE id = $4`)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %v", err)
	}
	defer stmt.Close()

	// Execute the statement
	if _, err := stmt.Exec(blog.Title, blog.Content, time.Now(), blog.ID); err != nil {
		return fmt.Errorf("failed to execute statement: %v", err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

func (m *BlogModels) DeleteBlog(id int) error {
	// Create a transaction
	tx, err := m.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	// Prepare the statement
	stmt, err := tx.Prepare(`DELETE FROM blogs WHERE id = $1`)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %v", err)
	}
	defer stmt.Close()

	// Execute the statement
	if _, err := stmt.Exec(id); err != nil {
		return fmt.Errorf("failed to execute statement: %v", err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

func (m *BlogModels) CreateComment(comment *forms.Comment) error {
	// Create a transaction
	tx, err := m.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	// Prepare the statement
	stmt, err := tx.Prepare(`INSERT INTO comments (blog_id, content) VALUES ($1, $2)`)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %v", err)
	}
	defer stmt.Close()

	// Execute the statement
	if _, err := stmt.Exec(comment.BlogID, comment.Content); err != nil {
		return fmt.Errorf("failed to execute statement: %v", err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

func (m *BlogModels) DeleteComment(id int) error {
	// Create a transaction
	tx, err := m.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	// Prepare the statement
	stmt, err := tx.Prepare(`DELETE FROM comments WHERE id = $1`)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %v", err)
	}
	defer stmt.Close()

	// Execute the statement
	if _, err := stmt.Exec(id); err != nil {
		return fmt.Errorf("failed to execute statement: %v", err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}
