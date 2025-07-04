package models

import (
	"fmt"
	"github.com/go-playground/validator/v10"

	"strings"
)

type Article struct {
	Id        int       `json:"id" db:"id"`
	UserId    int       `json:"user_id" db:"userId"`
	Title     string    `json:"title" db:"title"`
	Content   string    `json:"content" db:"content"`
	Likes     int       `json:"likes" db:"likes"`
	Comments  []Comment `json:"comments"`
	CreatedAt string    `json:"created_at" db:"created_at"`
	UpdatedAt string    `json:"updated_at" db:"updated_at"`
}

type ArticleRequest struct {
	Title   string `json:"title" validate:"required,min=3,max=40" msg:"Title must be between 3 and 40 characters"`
	Content string `json:"content" validate:"required,min=10,max=1000" msg:"Content must be between 10 and 1000 characters"`
}

func (a *ArticleRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(a)
	if err != nil {
		var sb strings.Builder
		for _, err := range err.(validator.ValidationErrors) {
			sb.WriteString(fmt.Sprintf("Field %s %s\n", err.Field(), err.Tag()))
		}
		return fmt.Errorf("%s", sb.String())
	}
	return nil
}
