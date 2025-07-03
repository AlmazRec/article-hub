package rest

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"restapp/internal/messages"
	"restapp/internal/models"
	"restapp/internal/response"
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
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: messages.ErrBadRequest,
			Error:   err.Error(),
		})
	}

	var req models.CommentRequest
	if err = c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: messages.ErrBadRequest,
			Error:   err.Error(),
		})
	}

	if err = req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: messages.ErrValidationFailed,
			Error:   err.Error(),
		})
	}

	claims, err := h.AuthService.ValidateToken(h.AuthService.FormatToken(c.Request().Header.Get("Authorization")))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: messages.ErrInvalidToken,
			Error:   err.Error(),
		})
	}

	if _, err := h.CommentService.CreateComment(ctx, &req, intArticleId, claims.UserId); err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: messages.ErrDatabaseOperation,
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response.SuccessResponse{
		Message: messages.MsgCommentCreated,
	})
}

func (h *CommentHandler) GetAllComments(c echo.Context) error {
	ctx := c.Request().Context()

	articleId := c.Param("id")
	intArticleId, err := strconv.Atoi(articleId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: messages.ErrBadRequest,
			Error:   err.Error(),
		})
	}

	comments, err := h.CommentService.GetAllComments(ctx, intArticleId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: messages.ErrDatabaseOperation,
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.SuccessResponse{
		Data: comments,
	})
}
