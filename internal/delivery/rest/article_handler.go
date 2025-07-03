package rest

import (
	"net/http"
	"restapp/internal/messages"
	"restapp/internal/models"
	"restapp/internal/response"
	"restapp/internal/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ArticleHandler struct {
	ArticleService services.ArticleServiceInterface
	AuthService    services.AuthServiceInterface
	CommentService services.CommentServiceInterface
}

func NewArticleHandler(articleService services.ArticleServiceInterface, authService services.AuthServiceInterface, commentService services.CommentServiceInterface) *ArticleHandler {
	return &ArticleHandler{ArticleService: articleService, AuthService: authService, CommentService: commentService}
}

func (h *ArticleHandler) GetAllArticles(c echo.Context) error {
	ctx := c.Request().Context()

	articles, err := (h.ArticleService).GetAllArticles(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: messages.ErrInternalServer,
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Data:    articles,
		Message: "",
	})
}

func (h *ArticleHandler) GetById(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: messages.ErrInvalidArticleID,
			Error:   err.Error(),
		})
	}

	article, err := h.ArticleService.GetById(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: messages.ErrArticleNotFound,
			Error:   err.Error(),
		})
	}

	comments, err := h.CommentService.GetAllComments(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: messages.ErrDatabaseOperation,
			Error:   err.Error(),
		})
	}

	article.Comments = *comments
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Data:    article,
		Message: "",
	})
}

func (h *ArticleHandler) StoreArticle(c echo.Context) error {
	ctx := c.Request().Context()

	var articleRequest models.ArticleRequest

	claims, err := h.AuthService.ValidateToken(h.AuthService.FormatToken(c.Request().Header.Get("Authorization")))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: messages.ErrInvalidToken,
			Error:   err.Error(),
		})
	}

	if err := c.Bind(&articleRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: messages.ErrInvalidArticleData,
			Error:   err.Error(),
		})
	}

	if err := articleRequest.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: messages.ErrValidationFailed,
			Error:   err.Error(),
		})
	}

	if err := h.ArticleService.CreateArticle(ctx, &articleRequest, claims.UserId); err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: messages.ErrDatabaseOperation,
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response.SuccessResponse{
		Message: messages.MsgArticleCreated,
	})
}

func (h *ArticleHandler) UpdateArticle(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: messages.ErrInvalidArticleID,
			Error:   err.Error(),
		})
	}

	var articleRequest models.ArticleRequest
	if err := c.Bind(&articleRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: messages.ErrInvalidArticleData,
			Error:   err.Error(),
		})
	}

	if err := articleRequest.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: messages.ErrValidationFailed,
			Error:   err.Error(),
		})
	}

	if err := h.ArticleService.UpdateArticle(ctx, id, &articleRequest); err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: messages.ErrArticleNotFound,
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.SuccessResponse{
		Message: messages.MsgArticleUpdated,
	})
}

func (h *ArticleHandler) DeleteArticle(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: messages.ErrInvalidArticleID,
			Error:   err.Error(),
		})
	}

	if err := h.ArticleService.DeleteArticle(ctx, id); err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: messages.ErrDatabaseOperation,
			Error:   err.Error(),
		})
	}

	return c.NoContent(http.StatusNoContent)
}
