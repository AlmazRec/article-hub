package models

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

type User struct {
	Id        int    `json:"id" db:"id"`
	Username  string `json:"username" db:"username"`
	Password  string `json:"password" db:"password"`
	Email     string `json:"email" db:"email"`
	Role      string `json:"role" db:"role"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Password string `json:"password" validate:"required,min=6"`
	Email    string `json:"email" validate:"required,email"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (r *RegisterRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)
	if err != nil {
		var sb strings.Builder
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				sb.WriteString(fmt.Sprintf("Field %s is required\n", err.Field()))
			case "email":
				sb.WriteString(fmt.Sprintf("Field %s must be a valid email address\n", err.Field()))
			case "min":
				sb.WriteString(fmt.Sprintf("Field %s must be at least %s characters long\n", err.Field(), err.Param()))
			case "max":
				sb.WriteString(fmt.Sprintf("Field %s must not exceed %s characters\n", err.Field(), err.Param()))
			default:
				sb.WriteString(fmt.Sprintf("Field %s failed validation: %s\n", err.Field(), err.Tag()))
			}
		}
		return fmt.Errorf("%s", sb.String())
	}
	return nil
}

func (l *LoginRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(l)
	if err != nil {
		var sb strings.Builder
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				sb.WriteString(fmt.Sprintf("Field %s is required\n", err.Field()))
			case "email":
				sb.WriteString(fmt.Sprintf("Field %s must be a valid email address\n", err.Field()))
			default:
				sb.WriteString(fmt.Sprintf("Field %s failed validation: %s\n", err.Field(), err.Tag()))
			}
		}
		return fmt.Errorf("%s", sb.String())
	}
	return nil
}
