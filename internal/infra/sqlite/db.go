package sqlite

import (
    "database/sql"
    "os"
)

import _ "github.com/mattn/go-sqlite3"

func Open(dbPath string) (*sql.DB, error) {
    if dbPath == "" {
        dbPath = "./app.db"
    }
    if err := os.MkdirAll("./", 0o755); err != nil {
        return nil, err
    }
    return sql.Open("sqlite3", dbPath)
}


