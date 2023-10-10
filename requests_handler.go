package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"

	"github.com/gorilla/sessions"
)

// Метод обработки запроса на регистрацию пользователя
func registerUserHandler(c echo.Context, db *sql.DB) error {
	var newUser User
	fmt.Println(newUser)
	if err := c.Bind(&newUser); err != nil {
		return c.String(http.StatusBadRequest, "Что-то не так...")
	}

	result, err := addUserToDB(db, newUser)
	fmt.Println(err)
	if err != nil {
		if strings.Contains(err.Error(), "уже существует") {
			return c.String(http.StatusBadRequest, "Пользователь с таким именем уже существует!")
		}
		return c.String(http.StatusInternalServerError, "Ошибка при сохранении пользователя в базе данных")
	}

	if err != nil {
		return c.String(http.StatusInternalServerError, "Ошибка при получении ID пользователя")
	}

	// Создаем сессию для пользователя, будем хранить id его записи в таблице
	sess, _ := session.Get("session", c)
	sess.Values["id"] = result

	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return c.String(http.StatusInternalServerError, "Ошибка сохранения сессии")
	}

	return c.String(http.StatusOK, "Вы успешно зарегистрированы!")
}

// Метод обработки запроса на авторизацию пользователя
func loginUserHandler(c echo.Context, db *sql.DB) error {
	var user User
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusBadRequest, "Что-то пошло не так")
	}

	dbUser, err := loginUserFromDB(db, user.NumberPhone, user.Password)
	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}

	sess, _ := session.Get("session", c)

	sess.Values["id"] = dbUser.Id

	sess.Save(c.Request(), c.Response())

	return c.String(http.StatusOK, "Вход выполнен успешно")
}

// Метод обработки запроса на получение данных о пользователе
// из базы данных, посредством кеширования в сессии id его записи
func getUserHandler(c echo.Context, db *sql.DB) error {
	// Получить сессию
	sess, _ := session.Get("session", c)
	userID, ok := sess.Values["id"].(int64)

	if !ok {
		// Если id пользователя отсутствует в сессии, вернуть ошибку
		return c.String(http.StatusUnauthorized, "Вы не авторизованы")
	}

	// Используйте userID для получения данных о пользователе из базы данных
	user, err := getUserInDB(db, userID)
	if err != nil {
		// Если произошла ошибка при запросе данных из базы данных, вернуть ошибку
		return c.String(http.StatusInternalServerError, "Ошибка при получении данных пользователя")
	}

	return c.JSON(http.StatusOK, user)
}

// Метод обработки запроса на выход пользователя из аккаунта
func logoutUserHandler(c echo.Context, db *sql.DB) error {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{MaxAge: -1}
	sess.Save(c.Request(), c.Response())

	return c.String(http.StatusOK, "Вы успешно вышли из аккаунта")
}
