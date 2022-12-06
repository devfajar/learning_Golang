package main

import (
	"net/http"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type User struct {
	Name  string `json:"name"  validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age"   validate:"gte=0,lte=80"`
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// Human Readable error
		if castedObject,  ok := err.(validator.ValidationErrors); ok {
			for _, err := range castedObject {
				switch err.Tag() {
				case "required":
					report.Message = fmt.Sprintf("%s is required", err.Field())
			case "email":
					report.Message = fmt.Sprintf("%s is not valid email", err.Field())
				case "gte":
					report.Message = fmt.Sprintf("%s value must be greater than %s", err.Field(), err.Param())
				case "lte":
					report.Message = fmt.Sprintf("%s value must be lower than %s", err.Field(), err.Param())
				}

				break
			}
		}

		c.Logger().Error(report)
		c.JSON(report.Code, report)
	}

	/* Custom Error Page
		e.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		errPage := fmt.Sprintf("%d.html", report.Code)
		if err := c.File(errPage); err != nil {
			c.HTML(report.Code, "Errrrooooorrrrr")
		}
	}

	e.GET("/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, true)
	}) */

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