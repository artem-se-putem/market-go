package main2

import (
    "database/sql"
    "fmt"
    "log"
    "time"
    
    _ "github.com/lib/pq"
)

func main2() {
    connStr := "user=postgres password=password dbname=postgres sslmode=disable"
    
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Failed to connect:", err)
    }
    defer db.Close()
    
    // Проверка подключения
    if err := db.Ping(); err != nil {
        log.Fatal("Ping failed:", err)
    }
    fmt.Println("✅ Successfully connected to PostgreSQL!")
    
    // Проверка версии
    var version string
    err = db.QueryRow("SELECT version()").Scan(&version)
    if err != nil {
        log.Fatal("Version check failed:", err)
    }
    fmt.Println("📋 PostgreSQL version:", version)
    
    // Проверка текущего времени БД
    var currentTime time.Time
    err = db.QueryRow("SELECT now()").Scan(&currentTime)
    if err != nil {
        log.Fatal("Time check failed:", err)
    }
    fmt.Println("⏰ Database time:", currentTime)
}