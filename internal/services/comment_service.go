package services

import (
	"context"
	"restapp/internal/models"
	"restapp/internal/repositories"
	"time"
)

type CommentServiceInterface interface {
	CreateComment(ctx context.Context, req *models.CommentRequest, articleId, userId int) (*models.Comment, error)
	GetAllComments(ctx context.Context, articleId int) (*[]models.Comment, error)
}

type CommentService struct {
	CommentRepository repositories.CommentRepositoryInterface
}

func NewCommentService(CommentRepository repositories.CommentRepositoryInterface) *CommentService {
	return &CommentService{
		CommentRepository: CommentRepository,
	}
}

func (s *CommentService) CreateComment(ctx context.Context, comment *models.CommentRequest, articleId, userId int) (*models.Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	commentModel := models.Comment{
		ArticleId: articleId,
		UserId:    userId,
		Content:   comment.Content,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	err := s.CommentRepository.CreateComment(ctx, &commentModel, articleId)
	if err != nil {
		return nil, err
	}
	return &commentModel, nil
}

func (s *CommentService) GetAllComments(ctx context.Context, articleId int) (*[]models.Comment, error) {
	comments, err := s.CommentRepository.GetAllComments(ctx, articleId)
	if err != nil {
		return nil, err
	}

	return comments, nil
}
