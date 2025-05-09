package models

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

type Comment struct {
	Id        int    `json:"id" db:"id"`
	ArticleId int    `json:"article_id" db:"article_id"`
	UserId    int    `json:"user_id" db:"user_id"`
	Content   string `json:"content" db:"content"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}

type CommentRequest struct {
	Content string `json:"content" validate:"required,min=10,max=1000" msg:"Content must be between 10 and 1000 characters"`
}

func (c *CommentRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		var sb strings.Builder
		for _, err := range err.(validator.ValidationErrors) {
			sb.WriteString(fmt.Sprintf("Field %s %s\n", err.Field(), err.Tag()))
		}
		return fmt.Errorf("%s", sb.String())
	}
	return nil
}
