package main

import "errors"

var (
	ErrUniqueviolation     = errors.New("not unique")
	ErrConstraintViolation = errors.New("Constraint")
	ErrForeighnKeyviol     = errors.New("Foreighnkey")
	ErrNorows              = errors.New("No rows")
)

type User struct {
	Name   string `json:"name" validate:"required" postgres:"name"`
	Age    int    `json:"age" validate:"max=120,min=19" postgres:"age"`
	Gender string `json:"gender" validate:"required,oneof=male female gay" postgres:"gender"`
}
