package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"restapp/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
)

type ArticleRepositoryInterface interface {
	GetAllArticles(ctx context.Context) (*[]models.Article, error)
	GetById(ctx context.Context, id int) (*models.Article, error)
	StoreArticle(ctx context.Context, article *models.Article, userId int) error
	UpdateArticle(ctx context.Context, id int, article *models.Article) error
	DeleteArticle(ctx context.Context, id int) error
}

type ArticleRepository struct {
	db *sqlx.DB
}

func NewArticleRepository(db *sqlx.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

func (r *ArticleRepository) GetAllArticles(ctx context.Context) (*[]models.Article, error) {
	var articles []models.Article
	err := r.db.Select(&articles, "SELECT id, title, content, created_at, updated_at FROM articles")
	if err != nil {
		return nil, fmt.Errorf("could not fetch articles: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return &articles, nil
}

func (r *ArticleRepository) GetById(ctx context.Context, id int) (*models.Article, error) {
	var article models.Article
	err := r.db.Get(&article, "SELECT id, title, content, created_at, updated_at FROM articles WHERE id = ?", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("article with id %d not found", id)
		}
		return nil, fmt.Errorf("could not fetch article: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return &article, nil
}

func (r *ArticleRepository) StoreArticle(ctx context.Context, article *models.Article, userId int) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := r.db.Exec("INSERT INTO articles (user_id, title, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", userId, article.Title, article.Content, article.CreatedAt, article.UpdatedAt)
	if err != nil {
		return fmt.Errorf("could not store article: %w", err)
	}

	return nil
}

func (r *ArticleRepository) UpdateArticle(ctx context.Context, id int, article *models.Article) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := r.db.Exec("UPDATE articles SET title = ?, content = ?, updated_at = ? WHERE id = ?", article.Title, article.Content, article.UpdatedAt, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("article with id %d not found", id)
		}
		return fmt.Errorf("could not update article: %w", err)
	}

	return nil
}

func (r *ArticleRepository) DeleteArticle(ctx context.Context, id int) error {
	_, err := r.db.Exec("DELETE FROM articles WHERE id =?", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("article with id %d not found", id)
		}
		return fmt.Errorf("could not delete article: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return nil
}
