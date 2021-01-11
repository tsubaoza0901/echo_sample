package main

import (
	"fmt"
	"net/http"

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
	e.POST("user", u.CreateUser)
	e.PUT("user/:id", u.UpdateUser)
	e.DELETE("user/:id", u.DeleteUser)
	e.GET("user/:id", u.GetUser)
	e.GET("users", u.GetUsers)
}

// --------
// handler↓
// --------

// CreateUser ...
func (u *User) CreateUser(c echo.Context) error {
	user := User{}

	if err := c.Bind(&user); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user.ID)
}

// UpdateUser ...
func (u *User) UpdateUser(c echo.Context) error {
	user := User{}

	if err := c.Bind(&user); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "Updated")
}

// DeleteUser ...
func (u *User) DeleteUser(c echo.Context) error {
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "Deleted")
}

// GetUser ...
func (u *User) GetUser(c echo.Context) error {
	user := User{}

	id := c.Param("id")

	fmt.Println(id, user)

	return c.JSON(http.StatusOK, id)
}

// GetUsers ...
func (u *User) GetUsers(c echo.Context) error {
	users := []*User{}

	fmt.Println(users)
	return c.JSON(http.StatusOK, users)
}

// --------
// main.go↓
// --------

func main() {
	e := echo.New()

	u := new(User)
	InitRouting(e, u)

	e.Start(":9005")
}
