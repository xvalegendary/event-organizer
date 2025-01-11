package main

import (
    "database/sql"
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers" // Для обработки CORS
    _ "github.com/mattn/go-sqlite3" // Импортируем драйвер SQLite
    "main/handler" 
)

func main() {
    // Создаем или открываем базу данных
    db, err := sql.Open("sqlite3", "./foo.db")
    if err != nil {
        fmt.Println("Ошибка при подключении к базе данных:", err)
        return
    }
    defer db.Close()

    // Создаем таблицу пользователей, если она не существует
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

    // Создаем маршрутизатор
    r := mux.NewRouter()
    r.HandleFunc("/auth", handler.AuthHandler).Methods("POST") // Определяем маршрут для аутентификации
    r.HandleFunc("/register", handler.RegisterHandler).Methods("POST") // Определяем маршрут для регистрации

    // Настройка CORS
    corsObj := handlers.AllowedOrigins([]string{"http://localhost:3000"}) // Разрешаем запросы с этого источника
    corsHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}) // Разрешаем заголовки
    corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}) // Разрешаем методы

    // Запускаем HTTP-сервер на порту 8000 с CORS
    fmt.Println("Сервер запущен на порту 8000")
    if err := http.ListenAndServe(":8000", handlers.CORS(corsObj, corsHeaders, corsMethods)(r)); err != nil {
        fmt.Println("Ошибка при запуске сервера:", err)
    }
}