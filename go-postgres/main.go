package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Blog struct {
	ID        int
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {
	db, err := sql.Open("pgx", "postgres://user:password@localhost:5432/blog")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	fmt.Println("Connected!")

	// Get all blogs
	rows, err := db.Query(
		`SELECT 
			b.id,
			b.title,
			b.content,
			b.created_at,
			b.updated_at
		FROM blogs AS b
		`,
	)
	if err != nil {
		log.Fatalf("failed to get all blogs: %v", err)
	}
	defer rows.Close()

	var blogs []Blog
	for rows.Next() {
		var blog Blog
		if err := rows.Scan(
			&blog.ID,
			&blog.Title,
			&blog.Content,
			&blog.CreatedAt,
			&blog.UpdatedAt,
		); err != nil {
			log.Fatalf("failed to scan blog: %v", err)
		}

		blogs = append(blogs, blog)
	}

	fmt.Printf("blogs: %v\n", blogs)
}
