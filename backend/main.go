package main

import (
    "database/sql"
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers" 
    _ "github.com/mattn/go-sqlite3" 
    "main/handler" 
)

func main() {
    
    db, err := sql.Open("sqlite3", "./foo.db")
    if err != nil {
        fmt.Println("Ошибка при подключении к базе данных:", err)
        return
    }
    defer db.Close()

    
    createTableSQL := `CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        login TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL,
        email TEXT NOT NULL UNIQUE
    );`
    if _, err := db.Exec(createTableSQL); err != nil {
        fmt.Println("Ошибка при создании таблицы:", err)
        return
    }

    
    r := mux.NewRouter()
    r.HandleFunc("/auth", handler.AuthHandler).Methods("POST") 
    r.HandleFunc("/register", handler.RegisterHandler).Methods("POST") 

   
    corsObj := handlers.AllowedOrigins([]string{"http://localhost:3000"}) 
    corsHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}) 
    corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}) 

   
    fmt.Println("Сервер запущен на порту 8000")
    if err := http.ListenAndServe(":8000", handlers.CORS(corsObj, corsHeaders, corsMethods)(r)); err != nil {
        fmt.Println("Ошибка при запуске сервера:", err)
    }
}