package main

import (
	"fmt"
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
	e.GET("use_session", u.UseSession)

	// e.POST("user", u.CreateUser)
	// e.PUT("user/:id", u.UpdateUser)
	// e.DELETE("user/:id", u.DeleteUser)
	// e.GET("user/:id", u.GetUser)
	// e.GET("users", u.GetUsers)
}

// --------
// handler↓
// --------

// Session ...
func Session(c echo.Context) error {
	sess, _ := session.Get("new_session_name", c)
	sess.ID = "session1"
	fmt.Println("sess.ID:", sess.ID) // => sess.ID:
	// fmt.Println("sess.IsNew:", sess.IsNew) // => sess.IsNew: true

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	// fmt.Println("sess.Options:", sess.Options) // => sess.Options: &{/  604800 false true 0}

	sess.Values["foo"] = "bar"
	// fmt.Println("sess.Values:", sess.Values) // => sess.Values: map[foo:bar]

	fmt.Println("sess:", sess)
	sess.Save(c.Request(), c.Response())
	return nil
}

// UseSession ...
func (u *User) UseSession(c echo.Context) error {

	Session(c)

	return c.JSON(http.StatusOK, "done")
}

// // CreateUser ...
// func (u *User) CreateUser(c echo.Context) error {
// 	user := User{}

// 	if err := c.Bind(&user); err != nil {
// 		return err
// 	}

// 	user = User{
// 		ID:   1,
// 		Name: user.Name,
// 		Age:  user.Age,
// 	}

// 	return c.JSON(http.StatusOK, user)
// }

// // UpdateUser ...
// func (u *User) UpdateUser(c echo.Context) error {
// 	user := User{}

// 	if err := c.Bind(&user); err != nil {
// 		return err
// 	}

// 	return c.JSON(http.StatusOK, "Updated")
// }

// // DeleteUser ...
// func (u *User) DeleteUser(c echo.Context) error {
// 	if err := c.Bind(u); err != nil {
// 		return err
// 	}
// 	return c.JSON(http.StatusOK, "Deleted")
// }

// // GetUser ...
// func (u *User) GetUser(c echo.Context) error {
// 	user := User{}

// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return err
// 	}

// 	// Getメソッドのイメージ
// 	if id == 1 {
// 		user = User{
// 			ID:   1,
// 			Name: "Tom",
// 			Age:  29,
// 		}
// 	} else if id == 2 {
// 		user = User{
// 			ID:   2,
// 			Name: "Bob",
// 			Age:  35,
// 		}

// 	} else {
// 		return c.JSON(http.StatusOK, "Not Found")
// 	}

// 	return c.JSON(http.StatusOK, user)
// }

// // GetUsers ...
// func (u *User) GetUsers(c echo.Context) error {
// 	users := []*User{}

// 	name := c.QueryParam("name")

// 	// Get Allメソッドのイメージ
// 	if name == "" {
// 		users = []*User{
// 			{
// 				ID:   1,
// 				Name: "Tom",
// 				Age:  29,
// 			},
// 			{
// 				ID:   2,
// 				Name: "Bob",
// 				Age:  35,
// 			},
// 		}
// 	} else if name == "Tom" {
// 		users = []*User{
// 			{
// 				ID:   1,
// 				Name: "Tom",
// 				Age:  29,
// 			},
// 		}
// 	} else if name == "Bob" {
// 		users = []*User{
// 			{
// 				ID:   2,
// 				Name: "Bob",
// 				Age:  35,
// 			},
// 		}
// 	} else {
// 		return c.JSON(http.StatusOK, "Not Found")
// 	}

// 	return c.JSON(http.StatusOK, users)
// }

// --------
// main.go↓
// --------

func main() {
	e := echo.New()

	// sessionを使用するMiddlewareとして追加
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	u := new(User)
	InitRouting(e, u)

	e.Start(":9005")
}
