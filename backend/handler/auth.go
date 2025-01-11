package handler

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "regexp"
    "golang.org/x/crypto/bcrypt"
    _ "github.com/mattn/go-sqlite3"
)

type AuthRequest struct {
    Login    string `json:"login"`
    Password string `json:"password"`
}

type RegisterRequest struct {
    Login    string `json:"login"`
    Password string `json:"password"`
    Email    string `json:"email"`
}

func isValidEmail(email string) bool {
    re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    return re.MatchString(email)
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
    var req AuthRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Неверный запрос", http.StatusBadRequest)
        return
    }

    // Подключаемся к базе данных SQLite
    db, err := sql.Open("sqlite3", "./foo.db")
    if err != nil {
        http.Error(w, "Не удалось подключиться к базе данных", http.StatusInternalServerError)
        return
    }
    defer db.Close() // Закрываем соединение с базой данных после завершения работы

    // Проверяем, существует ли логин в базе данных
    var hashedPassword string
    err = db.QueryRow("SELECT password FROM users WHERE login = ?", req.Login).Scan(&hashedPassword)
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "Неверный логин или пароль", http.StatusUnauthorized)
            return
        }
        http.Error(w, "Ошибка при проверке логина", http.StatusInternalServerError)
        return
    }

    // Сравниваем хешированный пароль с введенным паролем
    if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password)); err != nil {
        http.Error(w, "Неверный логин или пароль", http.StatusUnauthorized)
        return
    }

    // Если логин и пароль верны, возвращаем сообщение об успешной аутентификации
    response := map[string]string{"message": "Аутентификация успешна", "login": req.Login}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    var req RegisterRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Неверный запрос", http.StatusBadRequest)
        return
    }

    // Подключаемся к базе данных SQLite
    db, err := sql.Open("sqlite3", "./foo.db")
    if err != nil {
        http.Error(w, "Не удалось подключиться к базе данных", http.StatusInternalServerError)
        return
    }
    defer db.Close()

    // Проверяем, существует ли уже пользователь с таким логином или email
    var exists bool
    err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE login = ? OR email = ?)", req.Login, req.Email).Scan(&exists)
    if err != nil {
        http.Error(w, "Ошибка при проверке существования пользователя", http.StatusInternalServerError)
        return
    }
    if exists {
        http.Error(w, "Пользователь с таким логином или email уже существует", http.StatusConflict)
        return
    }

    // Хешируем пароль
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Ошибка при хешировании пароля", http.StatusInternalServerError)
        return
    }

    // Вставляем нового пользователя в базу данных
    _, err = db.Exec("INSERT INTO users (login, password, email) VALUES (?, ?, ?)", req.Login, hashedPassword, req.Email)
    if err != nil {
        http.Error(w, "Ошибка при регистрации пользователя", http.StatusInternalServerError)
        return
    }

    // Возвращаем сообщение об успешной регистрации
    response := map[string]string{"message": "Регистрация успешна", "login": req.Login}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}