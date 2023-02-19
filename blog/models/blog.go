package models

import (
	"blog/forms"
	"database/sql"
	"fmt"
)

type BlogModel struct {
	db *sql.DB
}

func NewBlogModel(db *sql.DB) *BlogModel {
	return &BlogModel{db: db}
}

func (m *BlogModel) GetAllBlogs() ([]forms.GetAllBlogsResponse, error) {
	rows, err := m.db.Query(
		`SELECT
			b.id,
			b.title,
			SUBSTRING(b.content FROM 0 FOR 350),
			b.created_at,
			b.updated_at,
			COUNT(c.id) AS comments
		FROM blogs AS b
		LEFT JOIN comments AS c ON c.blog_id = b.id
		GROUP BY b.id`,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get all blogs: %v", err)
	}
	defer rows.Close()

	var blogs []forms.GetAllBlogsResponse
	for rows.Next() {
		var blog forms.GetAllBlogsResponse
		if err := rows.Scan(
			&blog.ID,
			&blog.Title,
			&blog.Comments,
			&blog.CreatedAt,
			&blog.UpdatedAt,
			&blog.Comments,
		); err != nil {
			return nil, fmt.Errorf("failed to scan blog: %v", err)
		}

		blogs = append(blogs, blog)
	}

	return blogs, nil
}
