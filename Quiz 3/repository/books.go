package repository

import (
	"Quiz-3/structs"
	"database/sql"
	"time"
)
type BookRepository struct {
	DB *sql.DB
}

func GetAllBooks(db *sql.DB) (result []structs.Books, err error) {
	sql := `
        SELECT 
			b.id, b.title, b.description, b.image_url, b.release_year, 
			b.price, b.total_page, b.thickness, b.category_id, 
			b.created_at, b.created_by, b.modified_at, b.modified_by,
			c.id, c.name, c.created_at, c.created_by, c.modified_by, c.modified_at
		FROM books b
		JOIN categories c ON b.category_id = c.id
    `

	rows, err := db.Query(sql)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var book structs.Books
		var category structs.Categories

		err = rows.Scan(
			&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear,
			&book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID,
			&book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy,
			&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy, &category.ModifiedBy, &category.ModifiedAt, // Perbaikan urutan sesuai dengan SQL
		)
		if err != nil {
			return
		}

		book.Category = &category
		result = append(result, book)
	}

	return
}

func InsertBook(db *sql.DB, book structs.Books) (err error) {
    if book.TotalPage > 100 {
        book.Thickness = "tebal"
    } else {
        book.Thickness = "tipis"
    }
	now:=time.Now()
	sql := `
        INSERT INTO books (
            id, title, description, image_url, release_year, price, 
            total_page, thickness, category_id, created_at, created_by
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
    `

	_, err = db.Exec(sql,
		book.ID, book.Title, book.Description, book.ImageURL, book.ReleaseYear,
		book.Price, book.TotalPage, book.Thickness, book.CategoryID,
		now, book.CreatedBy,
	)
	return
}


func UpdateBook(db *sql.DB, book structs.Books) (err error) {
    if book.TotalPage > 100 {
        book.Thickness = "tebal"
    } else {
        book.Thickness = "tipis"
    }
    sql := `
        UPDATE books
        SET 
            title = $1, 
            description = $2, 
            image_url = $3, 
            release_year = $4, 
            price = $5, 
            total_page = $6, 
            thickness = $7, 
            category_id = $8, 
            modified_at = $9, 
            modified_by = $10
        WHERE id = $11
    `
	now := time.Now()
    _, err = db.Exec(sql,
        book.Title, book.Description, book.ImageURL, book.ReleaseYear,
        book.Price, book.TotalPage, book.Thickness, book.CategoryID,
        now, book.ModifiedBy, book.ID,
    )
    return
}


func DeleteBook(db *sql.DB, book structs.Books) (err error) {
    sql := "DELETE FROM books WHERE id = $1"

    errs := db.QueryRow(sql, book.ID)
    return errs.Err()
}
func GetBooksByCategoryID(db *sql.DB, categoryID int) ([]structs.Books, error) {
    query := `
        SELECT 
            b.id, b.title, b.description, b.image_url, b.release_year, b.price, b.total_page, b.thickness, b.category_id,
            b.created_at, b.created_by, b.modified_at, b.modified_by,
            c.id, c.name, c.created_by, c.created_at,c.modified_by, c.modified_at
        FROM 
            books b
        JOIN 
            categories c ON b.category_id = c.id
        WHERE 
            b.category_id = $1
    `

    rows, err := db.Query(query, categoryID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var books []structs.Books
    for rows.Next() {
        var book structs.Books
        var category structs.Categories

        err := rows.Scan(
            &book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear,
            &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID,
            &book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy,
            &category.ID, &category.Name, &category.CreatedBy, &category.CreatedAt, &category.ModifiedBy, &category.ModifiedAt,
        )
        if err != nil {
            return nil, err
        }

        book.Category = &category
        books = append(books, book)
    }

    return books, nil
}
