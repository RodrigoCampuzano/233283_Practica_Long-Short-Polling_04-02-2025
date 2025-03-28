package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLConnection(user, password, host, port, dbName string) (*sql.DB, error) {
    connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci", 
        user, password, host, port, dbName)
    
    db, err := sql.Open("mysql", connectionString)
    if err != nil {
        return nil, err
    }

    // Configuración de pool de conexiones
    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(25)
    db.SetConnMaxLifetime(5 * time.Minute)

    if err = db.Ping(); err != nil {
        return nil, err
    }

    return db, nil
}