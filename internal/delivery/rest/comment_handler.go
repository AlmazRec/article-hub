package rest

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"restapp/internal/errors"
	"restapp/internal/messages"
	"restapp/internal/models"
	"restapp/internal/services"
	"strconv"
)

type CommentHandler struct {
	CommentService services.CommentServiceInterface
	AuthService    services.AuthServiceInterface
}

func NewCommentHandler(commentService services.CommentServiceInterface, authService services.AuthServiceInterface) *CommentHandler {
	return &CommentHandler{
		CommentService: commentService,
		AuthService:    authService,
	}
}

func (h *CommentHandler) CreateComment(c echo.Context) error {
	ctx := c.Request().Context()

	articleId := c.Param("id")
	intArticleId, err := strconv.Atoi(articleId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.NewErrorResponse(
			http.StatusBadRequest,
			messages.ErrBadRequest,
			err.Error(),
		))
	}

	var req models.CommentRequest
	if err = c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, errors.NewErrorResponse(
			http.StatusBadRequest,
			messages.ErrBadRequest,
			err.Error(),
		))
	}

	if err = req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, errors.NewErrorResponse(
			http.StatusBadRequest,
			messages.ErrValidationFailed,
			err.Error(),
		))
	}

	claims, err := h.AuthService.ValidateToken(h.AuthService.FormatToken(c.Request().Header.Get("Authorization")))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.NewErrorResponse(
			http.StatusBadRequest,
			messages.ErrInvalidToken,
			err.Error(),
		))
	}

	if _, err := h.CommentService.CreateComment(ctx, &req, intArticleId, claims.UserId); err != nil {
		return c.JSON(http.StatusInternalServerError, errors.NewErrorResponse(
			http.StatusInternalServerError,
			messages.ErrDatabaseOperation,
			err.Error(),
		))
	}

	return c.JSON(http.StatusCreated, messages.MsgCommentCreated)
}

func (h *CommentHandler) GetAllComments(c echo.Context) error {
	ctx := c.Request().Context()

	articleId := c.Param("id")
	intArticleId, err := strconv.Atoi(articleId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.NewErrorResponse(
			http.StatusBadRequest,
			messages.ErrBadRequest,
			err.Error(),
		))
	}

	comments, err := h.CommentService.GetAllComments(ctx, intArticleId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.NewErrorResponse(
			http.StatusInternalServerError,
			messages.ErrDatabaseOperation,
			err.Error(),
		))
	}

	return c.JSON(http.StatusOK, comments)
}
