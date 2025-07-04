package services

import (
	"context"
	"fmt"
	"restapp/internal/messages"
	"restapp/internal/models"
	"restapp/internal/repositories"
	"time"
)

type ArticleServiceInterface interface {
	GetAllArticles(ctx context.Context) (*[]models.Article, error)
	GetById(ctx context.Context, id int) (*models.Article, error)
	CreateArticle(ctx context.Context, article *models.ArticleRequest, userId int) error
	UpdateArticle(ctx context.Context, id int, article *models.ArticleRequest) error
	DeleteArticle(ctx context.Context, id int) error
	LikeArticle(ctx context.Context, articleId int, userId int) error
	UnlikeArticle(ctx context.Context, articleId int, userId int) error
}

type ArticleService struct {
	r repositories.ArticleRepositoryInterface
}

func NewArticleService(r repositories.ArticleRepositoryInterface) *ArticleService {
	return &ArticleService{r: r}
}

func (s *ArticleService) GetAllArticles(ctx context.Context) (*[]models.Article, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	articles, err := s.r.GetAllArticles(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", messages.ErrGettingArticles, err)
	}

	return articles, nil
}

func (s *ArticleService) GetById(ctx context.Context, id int) (*models.Article, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	article, err := s.r.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", messages.ErrGettingArticles, err)
	}

	return article, nil
}

func (s *ArticleService) CreateArticle(ctx context.Context, article *models.ArticleRequest, userId int) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	articleModel := models.Article{
		Title:     article.Title,
		Content:   article.Content,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	return s.r.StoreArticle(ctx, &articleModel, userId)
}

func (s *ArticleService) UpdateArticle(ctx context.Context, id int, article *models.ArticleRequest) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	articleModel := models.Article{
		Id:        id,
		Title:     article.Title,
		Content:   article.Content,
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	return s.r.UpdateArticle(ctx, id, &articleModel)
}

func (s *ArticleService) DeleteArticle(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return s.r.DeleteArticle(ctx, id)
}

func (s *ArticleService) LikeArticle(ctx context.Context, articleId int, userId int) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return s.r.LikeArticle(ctx, articleId, userId)
}

func (s *ArticleService) UnlikeArticle(ctx context.Context, articleId int, userId int) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return s.r.UnlikeArticle(ctx, articleId, userId)
}
