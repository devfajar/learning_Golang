package main

import (
	"github.com/labstack/echo"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type CustomValidator struct {
	validator *validator.Validate
}

type User struct {
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age"   validate:"gte=0,lte=80"`
}


func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}


	// Routes here
	e.POST("/users", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}
		if err := c.Validate(u); err != nil {
			return err
		}

		return c.JSON(http.StatusOK, true)
	})
	e.Logger.Fatal(e.Start(":9000"))
}