package main

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// --------
// model↓
// --------

// User ...
type User struct {
	ID   uint   `json:"id" param:"id"`
	Name string `json:"name" validate:"required,yamada"` // validateタグ
	Age  int    `json:"age"`
	// SomethingArray []int  `json:"array" query:"array"`                   // queryタグ

}

// --------
// validator↓
// --------

// CustomValidator ...
type CustomValidator struct {
	validator *validator.Validate
}

// Validate ...
func (cv *CustomValidator) Validate(i interface{}) error {
	cv.validator.RegisterValidation("yamada", CustomValidate)
	return cv.validator.Struct(i)
}

// CustomValidate ...
func CustomValidate(fl validator.FieldLevel) bool {
	return fl.Field().String() == "yamada"
}

// NewCustomValidator ...
func NewCustomValidator() echo.Validator {
	return &CustomValidator{validator: validator.New()}
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
		log.Error(err)
		return err
	}
	// バリデーションの呼び出し
	if err := c.Validate(&user); err != nil {
		log.Error(err)
		return err
	}

	user = User{
		ID:   1,
		Name: user.Name,
		Age:  user.Age,
	}

	return c.JSON(http.StatusOK, user)
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

	if err := c.Bind(&user); err != nil {
		return err
	}

	// fmt.Println(user.ID)
	// fmt.Println(user.Name)
	// fmt.Println(user.Age)
	// fmt.Println(user.SomethingArray)
	// for _, v := range user.SomethingArray {
	// 	fmt.Println(v)
	// }

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	// Getメソッドのイメージ
	if id == 1 {
		user = User{
			ID:   1,
			Name: "Tom",
			Age:  29,
		}
	} else if id == 2 {
		user = User{
			ID:   2,
			Name: "Bob",
			Age:  35,
		}

	} else {
		return c.JSON(http.StatusOK, "Not Found")
	}

	return c.JSON(http.StatusOK, user)
}

// GetUsers ...
func (u *User) GetUsers(c echo.Context) error {
	users := []*User{}

	name := c.QueryParam("name")

	// Get Allメソッドのイメージ
	if name == "" {
		users = []*User{
			{
				ID:   1,
				Name: "Tom",
				Age:  29,
			},
			{
				ID:   2,
				Name: "Bob",
				Age:  35,
			},
		}
	} else if name == "Tom" {
		users = []*User{
			{
				ID:   1,
				Name: "Tom",
				Age:  29,
			},
		}
	} else if name == "Bob" {
		users = []*User{
			{
				ID:   2,
				Name: "Bob",
				Age:  35,
			},
		}
	} else {
		return c.JSON(http.StatusOK, "Not Found")
	}

	return c.JSON(http.StatusOK, users)
}

// --------
// main.go↓
// --------

func main() {
	e := echo.New()

	// CustomValidatorのインスタンス生成
	e.Validator = NewCustomValidator()

	u := new(User)
	InitRouting(e, u)

	e.Start(":9005")
}
