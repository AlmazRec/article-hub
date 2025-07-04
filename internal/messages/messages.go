package messages

import (
	"errors"
)

var (
	// Auth messages
	ErrInvalidCredentials    = errors.New("invalid email or password")
	ErrEmailAlreadyExists    = errors.New("email already exists")
	ErrUsernameAlreadyExists = errors.New("username already exists")
	ErrInvalidToken          = errors.New("invalid token")
	ErrMissingToken          = errors.New("missing authorization token")
	ErrParsingToken          = errors.New("error parsing token")
	ErrTokenExpired          = errors.New("token expired")
	ErrInvalidSigningMethod  = errors.New("invalid signing method")
	ErrSigningToken          = errors.New("error signing token")
	ErrHashingPassword       = errors.New("error hashing password")
	ErrBadRequest            = errors.New("bad request")
	ErrConvertingExpTime     = errors.New("error converting expiration time")
	ErrComparingPasswords    = errors.New("error comparing passwords")
	ErrCreatingUser          = errors.New("error creating user")
	ErrGeneratingToken       = errors.New("error generating token")

	// Article messages
	ErrFetchArticles      = errors.New("failed to fetch articles")
	ErrArticleNotFound    = errors.New("article not found")
	ErrInvalidArticleData = errors.New("invalid article data")
	ErrInvalidArticleID   = errors.New("invalid article ID")
	ErrGettingArticles    = errors.New("error getting articles")
	ErrLikeExists         = errors.New("like already exists")
	MsgArticleCreated     = "article successfully created"
	MsgArticleUpdated     = "article successfully updated"
	MsgArticleDeleted     = "article successfully deleted"
	MsgArticleLiked       = "article successfully liked"
	MsgArticleUnliked     = "article successfully unliked"

	// Validation messages
	ErrValidationFailed   = errors.New("validation failed")
	ErrFieldRequired      = "field %s is required"
	ErrFieldMinLength     = "field %s must be at least %d characters"
	ErrFieldMaxLength     = "field %s must not exceed %d characters"
	ErrInvalidEmailFormat = "invalid email format"

	// Success messages
	MsgRegistrationSuccess = "registration successful"
	MsgLoginSuccess        = "login successful"

	// Server errors
	ErrInternalServer     = errors.New("internal server error")
	ErrDatabaseConnection = errors.New("database connection error")
	ErrDatabaseOperation  = errors.New("database operation failed")

	// Comment messages
	MsgCommentCreated  = "comment successfully created"
	MsgCommentUpdated  = "comment successfully updated"
	ErrGettingComments = errors.New("error to getting comments")
	ErrCreatingComment = errors.New("error creating comment")

	// User messages
	ErrGettingUser = errors.New("error getting user")
)
