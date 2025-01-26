package repository

import (
	"Quiz-3/structs"
	"database/sql"
	"time"
)

func GetAllCategories(db *sql.DB) (result []structs.Categories, err error) {
    sql := "SELECT * FROM categories"

    rows, err := db.Query(sql)
    if err != nil {
       return
    }

    defer rows.Close()
    for rows.Next() {
       var category structs.Categories

       err = rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy, &category.ModifiedAt, &category.ModifiedBy)
       if err != nil {
          return
       }

       result = append(result, category)
    }

    return
}

func InsertCategory(db *sql.DB, category structs.Categories) error {
    sql := "INSERT INTO categories(id, name, created_at, created_by) VALUES ($1, $2, $3, $4)"
    now := time.Now()
    _, err := db.Exec(sql, category.ID, category.Name, now, category.CreatedBy)


    return err
}


func UpdateCategory(db *sql.DB, category structs.Categories) error {
    sql := "UPDATE categories SET name = $1, modified_at = $2, modified_by = $3 WHERE id = $4"

    now := time.Now()

    _, err := db.Exec(sql, category.Name, now, category.ModifiedBy, category.ID)
    return err
}


func DeleteCategory(db *sql.DB, category structs.Categories) (err error) {
    sql := "DELETE FROM categories WHERE id = $1"

    errs := db.QueryRow(sql, category.ID)
    return errs.Err()
}

