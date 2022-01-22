package main

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// --------
// model↓
// --------

// User ...
type User struct {
    ID   uint   `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}

// --------
// router↓
// --------

// InitRouting ...
func InitRouting(e *echo.Echo, u *User) {
    e.GET("login", u.Login)
	e.GET("users", u.GetUsers)
}

// --------
// middleware↓
// --------

func Middleware(e *echo.Echo) {
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
}

// --------
// handler↓
// --------

// Login ...
func (u *User) Login(c echo.Context) error {
	// get new session
    sess, err := session.Get("session_name", c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to session.Get()")
	}

	// if !sess.IsNew {
	// 	// session発行済みの場合はログイン済みとしてトップページにリダイレクト
	// 	return c.Redirect(http.StatusOK, "/top_page")
	// }

	// set options
    sess.Options = &sessions.Options{
        Path:     "/",
        MaxAge:   86400 * 7,
        HttpOnly: true,
    }

	// set values
	sess.Values["token"] = "xxx"

	// save session
    if err := sess.Save(c.Request(), c.Response()); err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to sess.Save()")
	}

    return c.Redirect(http.StatusFound, "users")
}

func(u *User) GetUsers(c echo.Context) error {
	sess, err := session.Get("session_name", c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to session.Get()")
	}
	if sess.IsNew {
		log.Printf("%+v", *sess)
		return c.JSON(http.StatusOK, "Can't Get Users, because session is new")
	}

	user := User{
		ID: 1,
		Name: "Yamada",
		Age: 30,
	}
	return c.JSON(http.StatusOK, user)
}

// --------
// main.go↓
// --------

func main() {
    e := echo.New()

    Middleware(e)

    u := new(User)
    InitRouting(e, u)

    e.Start(":9005")
}