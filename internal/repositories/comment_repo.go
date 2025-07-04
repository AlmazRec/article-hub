package repositories

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"restapp/internal/messages"
	"restapp/internal/models"
	"time"
)

type CommentRepositoryInterface interface {
	CreateComment(ctx context.Context, comment *models.Comment, articleId int) error
	GetAllComments(ctx context.Context, articleId int) ([]models.Comment, error)
}

type CommentRepository struct {
	db *sqlx.DB
}

func NewCommentRepository(db *sqlx.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) CreateComment(ctx context.Context, comment *models.Comment, articleId int) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(
		ctx,
		`INSERT INTO comments (article_id, user_id, content, created_at, updated_at)
		 VALUES (?, ?, ?, ?, ?)`,
		articleId,
		comment.UserId,
		comment.Content,
		comment.CreatedAt,
		comment.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("%w: %v", messages.ErrCreatingComment, err)
	}

	return nil
}

func (r *CommentRepository) GetAllComments(ctx context.Context, articleId int) ([]models.Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var comments []models.Comment

	err := r.db.SelectContext(
		ctx,
		&comments,
		`SELECT id, article_id, user_id, content, created_at, updated_at
		 FROM comments
		 WHERE article_id = ?`,
		articleId,
	)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", messages.ErrGettingComments, err)
	}

	return comments, nil
}
