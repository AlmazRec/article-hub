package repositories

import (
	"context"
	"database/sql"
	"restapp/internal/messages"
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
	LikeArticle(ctx context.Context, articleId int, userId int) error
	UnlikeArticle(ctx context.Context, articleId int, userId int) error
}

type ArticleRepository struct {
	db *sqlx.DB
}

func NewArticleRepository(db *sqlx.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

func (r *ArticleRepository) GetAllArticles(ctx context.Context) (*[]models.Article, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var articles []models.Article
	err := r.db.SelectContext(
		ctx,
		&articles,
		`SELECT a.id, a.title, a.content, COUNT(l.id) as likes, a.created_at, a.updated_at
			FROM 
			    articles a 
			LEFT JOIN 
				likes l ON a.id = l.article_id
			GROUP BY 
    			a.id, a.title, a.content, a.created_at, a.updated_at;`,
	)
	if err != nil {
		return nil, messages.ErrFetchArticles
	}
	return &articles, nil
}

func (r *ArticleRepository) GetById(ctx context.Context, id int) (*models.Article, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var article models.Article
	err := r.db.GetContext(ctx,
		&article,
		`SELECT id, title, content, created_at, updated_at FROM articles WHERE id = ?`,
		id,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, messages.ErrArticleNotFound
		}
		return nil, messages.ErrFetchArticles
	}
	return &article, nil
}

func (r *ArticleRepository) StoreArticle(ctx context.Context, article *models.Article, userId int) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(
		ctx,
		`INSERT INTO articles (user_id, title, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`,
		userId,
		article.Title,
		article.Content,
		article.CreatedAt,
		article.UpdatedAt,
	)
	if err != nil {
		return messages.ErrInvalidArticleData
	}
	return nil
}

func (r *ArticleRepository) UpdateArticle(ctx context.Context, id int, article *models.Article) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	result, err := r.db.ExecContext(
		ctx,
		`UPDATE articles
		SET title = ?, content = ?, updated_at = ?
		WHERE id = ?`,
		article.Title,
		article.Content,
		article.UpdatedAt,
		id,
	)
	if err != nil {
		return messages.ErrFetchArticles
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return messages.ErrDatabaseOperation
	}
	if rowsAffected == 0 {
		return messages.ErrArticleNotFound
	}
	return nil
}

func (r *ArticleRepository) DeleteArticle(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	result, err := r.db.ExecContext(
		ctx,
		`DELETE FROM articles
       	WHERE id = ?`,
		id,
	)
	if err != nil {
		return messages.ErrFetchArticles
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return messages.ErrDatabaseOperation
	}
	if rowsAffected == 0 {
		return messages.ErrArticleNotFound
	}
	return nil
}

func (r *ArticleRepository) LikeArticle(ctx context.Context, articleId int, userId int) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var exists bool
	err := r.db.GetContext(
		ctx,
		&exists,
		`SELECT EXISTS (
            SELECT 1 FROM likes WHERE article_id = ? AND user_id = ?
        )`,
		articleId,
		userId,
	)
	if err != nil {
		return messages.ErrDatabaseOperation
	}
	if exists {
		return messages.ErrLikeExists
	}

	_, err = r.db.ExecContext(
		ctx,
		`INSERT INTO likes (article_id, user_id) 
				VALUES (?, ?);`,
		articleId,
		userId,
	)

	if err != nil {
		return messages.ErrDatabaseOperation
	}
	return nil
}

func (r *ArticleRepository) UnlikeArticle(ctx context.Context, articleId int, userId int) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(
		ctx,
		`DELETE FROM likes WHERE article_id = ? AND user_id = ?`,
		articleId,
		userId,
	)
	if err != nil {
		return messages.ErrDatabaseOperation
	}

	return nil
}
