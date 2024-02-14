package db

import (
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
)

func Connect(dsn string) (*sql.DB, error) {
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Println("Error connecting to the database:", err)
        return nil, err
    }
    return db, nil
}