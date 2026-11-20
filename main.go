package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
	"uuid"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
)

type user struct {
	Name   string `json:"name" validate:"required"`
	Age    int    `json:"age" validate:"max=120,min=19"`
	Gender string `json:"gender" validate:"required,oneof=male female gay"`
}
type customvalidator struct {
	valdator *validator.Validate
}

func (c *customvalidator) Validate(i any) error {
	if err := c.valdator.Struct(i); err != nil {
		var validationerr validator.ValidationErrors
		if errors.As(err, &validationerr) {
			for _, v := range validationerr {
				fmt.Println(v.Tag(), v.Field(), v.Param())
			}
			return echo.NewHTTPError(http.StatusBadRequest, "errors fmted")
		}
		return echo.NewHTTPError(http.StatusBadRequest, "fuck nigga")
	}
	return nil
}

func hellohandler(c *echo.Context) error {
	var usr user
	err := c.Bind(&usr)
	if err != nil {
		return c.String(http.StatusBadRequest, "bind err")
	}
	if err = c.Validate(usr); err != nil {
		return c.String(http.StatusBadRequest, "validater err")
	}
	return c.JSON(http.StatusOK, usr)
}

func idsetter(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c *echo.Context) error {
		id := uuid.New().String()
		fmt.Println(id)
		v := c.Request().Header.Get("token")
		if v != "nar" {
			return c.JSON(http.StatusBadRequest, "token isnt right")
		}
		st := time.Now()
		res := next(c)
		end := time.Since(st)
		fmt.Println("it took:", end)
		return res
	}

}

func main() {
	e := echo.New()
	e.Validator = &customvalidator{validator.New()}
	mainn := e.Group("/main")
	mainn.Use(idsetter)
	mainn.GET("/hello", hellohandler)
	if err := e.Start(":8080"); err != nil {
		log.Fatal("error start")
	}
}
