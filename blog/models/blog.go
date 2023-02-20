package models

import (
	"database/sql"
	"fmt"

	"blog/forms"
)

// BLogModel wraps a sql.DB connection pool.
type BlogModel struct {
	db *sql.DB
}

// NewBlogModel returns a new BlogModel.
func NewBlogModel(db *sql.DB) *BlogModel {
	return &BlogModel{db: db}
}

// CreateBlog inserts a new blog into the database.
func (m *BlogModel) CreateBlog(blog forms.CreateBlogRequest) error {
	// Prepare the SQL statement. This returns a sql.Stmt object, which can be
	// used to execute the statement multiple times with different data.
	//
	// The $1 and $2 placeholders are used to represent the title and content
	// parameters. This is known as "SQL parameter binding". It's a good idea to
	// use parameter binding to prevent SQL injection attacks.
	//
	// For more information on SQL parameter binding, see:
	// https://www.calhoun.io/inserting-records-into-a-postgresql-database-with-gos-database-sql-package/
	stmt, err := m.db.Prepare(`
		INSERT INTO blogs (title, content)
		VALUES ($1, $2)
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %v", err)
	}
	defer stmt.Close() // Remember to close the statement when you're done with it!

	// Execute the statement, passing in the title and content parameters.
	//
	// The Exec method returns a sql.Result object, which contains some basic
	// information about what happened when the statement was executed.
	//
	// We don't actually need the sql.Result object in this case, so we simply
	// discard it with the blank identifier.
	//
	// If you want to get the ID of the newly inserted blog, you can use the
	// LastInsertId method on the sql.Result object.
	//
	// For more information on the LastInsertId method, see:
	// https://golang.org/pkg/database/sql/#Result
	//
	// If you want to get the number of rows affected by the statement, you can
	// use the RowsAffected method on the sql.Result object.
	//
	// For more information on the RowsAffected method, see:
	// https://golang.org/pkg/database/sql/#Result
	//
	// If you want to get both the ID of the newly inserted blog and the number
	// of rows affected by the statement, you can use the Exec method on the
	// sql.DB object instead of the sql.Stmt object.
	//
	// For more information on the Exec method, see:
	// https://golang.org/pkg/database/sql/#DB
	if _, err := stmt.Exec(blog.Title, blog.Content); err != nil {
		return fmt.Errorf("failed to execute statement: %v", err)
	}

	return nil
}

// GetAllBlogs returns all blogs from the database.
func (m *BlogModel) GetAllBlogs() ([]forms.GetAllBlogsResponse, error) {
	// We use the LEFT JOIN to ensure that we get a row for every blog, even if
	// it doesn't have any comments. This is so that we can display the number of
	// comments for each blog.
	//
	// We use the SUBSTRING function to limit the length of the blog content to
	// 350 characters. This is so that we can display a short preview of the blog
	// content on the home page.
	//
	// We use the COUNT function to count the number of comments for each blog.
	// This is so that we can display the number of comments for each blog on the
	// home page.
	//
	// We use the GROUP BY clause to group the results by blog ID. This is so
	// that we don't get duplicate rows in the result set.
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
	defer rows.Close() // Remember to close the rows when you're done with them!

	// Initialize an empty slice to hold the blogs.
	var blogs []forms.GetAllBlogsResponse
	for rows.Next() {
		// Initialize a new blog struct.
		var blog forms.GetAllBlogsResponse
		// Use rows.Scan to copy the values from each field in the row into the
		// corresponding field in the blog struct.
		//
		// Notice that we pass a pointer to the blog.ID field as the first
		// argument. This is because Scan requires a pointer to the value that it
		// copies the data into.
		//
		// If we didn't pass a pointer, then Scan would try to copy the data into
		// a copy of the blog.ID field, and the changes would be lost once Scan
		// returned.
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

		blogs = append(blogs, blog)
	}

	return blogs, nil
}

// GetBlogByID returns a single blog from the database, based on its ID.
func (m *BlogModel) GetBlogByID(id int) (forms.GetBlogByIDResponse, error) {
	// Prepare the SQL statement. This returns a sql.Stmt object, which can be
	// used to execute the statement multiple times with different data.
	stmt, err := m.db.Prepare(`
		SELECT
			id,
			title,
			content,
			created_at,
			updated_at
		FROM blogs
		WHERE id = $1
	`)
	if err != nil {
		return forms.GetBlogByIDResponse{}, fmt.Errorf("failed to prepare statement: %v", err)
	}
	defer stmt.Close() // Remember to close the statement when you're done with it!

	// Execute the statement, passing in the id parameter.
	//
	// The QueryRow method returns a sql.Row object, which can be used to scan
	// the result set into a struct.
	//
	// For more information on the QueryRow method, see:
	// https://golang.org/pkg/database/sql/#DB
	row := stmt.QueryRow(id)

	// Initialize a new blog struct.
	var blog forms.GetBlogByIDResponse
	// Use row.Scan to copy the values from each field in the row into the
	// corresponding field in the blog struct.
	if err := row.Scan(
		&blog.ID,
		&blog.Title,
		&blog.Content,
		&blog.CreatedAt,
		&blog.UpdatedAt,
	); err != nil {
		return forms.GetBlogByIDResponse{}, fmt.Errorf("failed to scan blog: %v", err)
	}

	// Get the comments for the blog.
	commentRows, err := m.db.Query(`
		SELECT
			id,
			blog_id,
			content,
			created_at,
			updated_at
		FROM comments
		WHERE blog_id = $1
	`, id)
	if err != nil {
		return forms.GetBlogByIDResponse{}, fmt.Errorf("failed to get comments: %v", err)
	}
	defer commentRows.Close() // Remember to close the rows when you're done with them!

	for commentRows.Next() {
		// Initialize a new comment struct.
		var comment forms.Comment
		if err := commentRows.Scan(
			&comment.ID,
			&comment.BlogID,
			&comment.Content,
			&comment.CreatedAt,
			&comment.UpdatedAt,
		); err != nil {
			return forms.GetBlogByIDResponse{}, fmt.Errorf("failed to scan comment: %v", err)
		}

		blog.Comments = append(blog.Comments, comment)
	}

	return blog, nil
}

// UpdateBlog updates a single blog in the database, based on its ID.
func (m *BlogModel) UpdateBlog(id int, blog forms.UpdateBlogRequest) error {
	// Prepare the SQL statement.
	//
	// The $1, $2 and $3 placeholders are used to represent the title, content
	// and id parameters.
	stmt, err := m.db.Prepare(`
		UPDATE blogs
		SET title = $1, content = $2
		WHERE id = $3
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %v", err)
	}
	defer stmt.Close() // Remember to close the statement when you're done with it!

	// Execute the statement, passing in the title, content and id parameters.
	if _, err := stmt.Exec(blog.Title, blog.Content, id); err != nil {
		return fmt.Errorf("failed to execute statement: %v", err)
	}

	return nil
}

// DeleteBlog deletes a single blog from the database, based on its ID.
func (m *BlogModel) DeleteBlog(id int) error {
	// Prepare the SQL statement.
	stmt, err := m.db.Prepare(`
		DELETE FROM blogs
		WHERE id = $1
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %v", err)
	}
	defer stmt.Close() // Remember to close the statement when you're done with it!

	// Execute the statement, passing in the id parameter.
	if _, err := stmt.Exec(id); err != nil {
		return fmt.Errorf("failed to execute statement: %v", err)
	}

	return nil
}

// CreateComment inserts a new comment into the database.
func (m *BlogModel) CreateComment(blogID int, comment forms.CreateCommentRequest) error {
	// Prepare the SQL statement.
	stmt, err := m.db.Prepare(`
		INSERT INTO comments (blog_id, content)
		VALUES ($1, $2)
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %v", err)
	}
	defer stmt.Close() // Remember to close the statement when you're done with it!

	// Execute the statement, passing in the blog_id and content parameters.
	if _, err := stmt.Exec(blogID, comment.Content); err != nil {
		return fmt.Errorf("failed to execute statement: %v", err)
	}

	return nil
}
