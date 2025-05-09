package repositories

import (
	"context"
	"github.com/jmoiron/sqlx"
	"restapp/internal/models"
	"time"
)

type CommentRepositoryInterface interface {
	CreateComment(ctx context.Context, comment *models.Comment, articleId int) error
	GetAllComments(ctx context.Context, articleId int) (*[]models.Comment, error)
}

type CommentRepository struct {
	db *sqlx.DB
}

func NewCommentRepository(db *sqlx.DB) *CommentRepository {
	return &CommentRepository{
		db: db,
	}
}

func (c *CommentRepository) CreateComment(ctx context.Context, comment *models.Comment, articleId int) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `INSERT INTO comments (article_id, user_id, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`
	_, err := c.db.ExecContext(ctx, query, articleId, comment.UserId, comment.Content, comment.CreatedAt, comment.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (c *CommentRepository) GetAllComments(ctx context.Context, articleId int) (*[]models.Comment, error) {
	var comments []models.Comment
	err := c.db.SelectContext(ctx, &comments, "SELECT id, article_id, user_id, content, created_at, updated_at FROM comments WHERE article_id = ?", articleId)
	if err != nil {
		return nil, err
	}
	return &comments, nil
}
