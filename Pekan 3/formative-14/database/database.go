package database

import (
    "database/sql"
    "embed"
    "fmt"
    migrate "github.com/rubenv/sql-migrate"
)

// Embed folder sql_migrations
//go:embed sql_migrations/*
var dbMigrations embed.FS

var DbConnection *sql.DB

func DBMigrate(dbParam *sql.DB) {
    migrations := &migrate.EmbedFileSystemMigrationSource{
        FileSystem: dbMigrations,
        Root:       "sql_migrations", // Pastikan path sesuai
    }

    n, errs := migrate.Exec(dbParam, "postgres", migrations, migrate.Up)
    if errs != nil {
        fmt.Printf("Migration failed: %v\n", errs)
        panic(errs)
    }

    DbConnection = dbParam

    fmt.Println("Migration success, applied", n, "migrations!")
}
