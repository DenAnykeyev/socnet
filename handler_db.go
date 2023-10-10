package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Метод для инициализации подключения к базе данных
func initDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/socnet")
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Метод для добавления записи пользователя в таблицу users
func addUserToDB(db *sql.DB, newUser User) (int64, error) {
	var existingUser User
	err := db.QueryRow("SELECT * FROM users WHERE numberPhone=?", newUser.NumberPhone).Scan(&existingUser.FirstName, &existingUser.LastName, &existingUser.NumberPhone, &existingUser.Password, &existingUser.Rules)

	if err == nil {
		return -1, fmt.Errorf("Пользователь с именем %s уже существует", newUser.FirstName)
	} else if err != sql.ErrNoRows {
		return 0, err
	}

	result, err := db.Exec("INSERT INTO users (firstName, lastName, numberPhone, password, rules) VALUES (?, ?, ?, ?, ?)",
		newUser.FirstName, newUser.LastName, newUser.NumberPhone, newUser.Password, newUser.Rules)
	if err != nil {
		return 0, err
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func loginUserFromDB(db *sql.DB, numberPhone, password string) (*User, error) {
	var user User
	err := db.QueryRow("SELECT id FROM users WHERE numberPhone = ? AND password = ?", numberPhone, password).Scan(&user.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Неверное имя пользователя или пароль")
		}
		return nil, fmt.Errorf("Ошибка при выполнении запроса")
	}
	return &user, nil
}

// Метод для получения записи пользователя по его
// id записи в хеше сессии
func getUserInDB(db *sql.DB, userID int64) (*User, error) {
	var user User
	err := db.QueryRow("SELECT id, firstName, lastName, numberPhone, rules FROM users WHERE id=?", userID).Scan(&user.Id, &user.FirstName, &user.LastName, &user.NumberPhone, &user.Rules)

	if err == sql.ErrNoRows {
		// Если запись не найдена, устанавливаем Id в -1
		user.Id = -1
		return &user, nil
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}
