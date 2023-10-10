package main

import (
	"fmt"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id          int    `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	NumberPhone string `json:"numberPhone"`
	Password    string `json:"password"`
	Rules       string `json:"rules"`
}

func main() {
	fmt.Println("\n[INFO] Проект запускается...")
	e := echo.New()

	e.Static("/public", "public")
	e.Static("/assets", "public/assets")

	// Инциализация соединения с базой данных
	db, err := initDB()
	if err != nil {
		fmt.Println(err)
		fmt.Println("[ERROR] Ошибка инициализирования подключения к базе данных!")
		return
	} else {
		fmt.Println("[INFO] Инциализация подключения к базе данных прошла успешно!")
	}
	defer db.Close()

	// Инициализация поддержки gorilla сессий для Echo
	store := sessions.NewCookieStore([]byte("secret_key"))
	e.Use(session.Middleware(store))

	e.POST("/register_user", func(c echo.Context) error {
		return registerUserHandler(c, db)
	})

	e.GET("/get_user", func(c echo.Context) error {
		return getUserHandler(c, db)
	})

	e.POST("/logout_user", func(c echo.Context) error {
		return logoutUserHandler(c, db)
	})

	e.GET("*", func(c echo.Context) error {
		return c.File("index.html")
	})

	e.Logger.Fatal(e.Start(":1323"))

}
