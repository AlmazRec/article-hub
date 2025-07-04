package messages

import "fmt"

var (
	// Auth messages
	ErrInvalidCredentials    = fmt.Errorf("invalid email or password")
	ErrEmailAlreadyExists    = fmt.Errorf("email already exists")
	ErrUsernameAlreadyExists = fmt.Errorf("username already exists")
	ErrInvalidToken          = fmt.Errorf("invalid token")
	ErrMissingToken          = fmt.Errorf("missing authorization token")
	ErrParsingToken          = fmt.Errorf("error parsing token: %v")
	ErrTokenExpired          = fmt.Errorf("token expired")
	ErrInvalidSigningMethod  = fmt.Errorf("invalid signing method")
	ErrSigningToken          = fmt.Errorf("error signing token: %v")
	ErrHashingPassword       = fmt.Errorf("error hashing password: %v")
	ErrBadRequest            = fmt.Errorf("bad request")
	ErrConvertingExpTime     = fmt.Errorf("error converting expiration time: %v")
	ErrComparingPasswords    = fmt.Errorf("error comparing passwords: %v")
	ErrCreatingUser          = fmt.Errorf("error creating user: %v")
	ErrGeneratingToken       = fmt.Errorf("error generating token: %v")

	// Article messages
	ErrFetchArticles      = fmt.Errorf("failed to fetch articles: %v")
	ErrArticleNotFound    = fmt.Errorf("article not found")
	ErrInvalidArticleData = fmt.Errorf("invalid article data")
	ErrInvalidArticleID   = fmt.Errorf("invalid article ID")
	ErrGettingArticles    = fmt.Errorf("error getting articles: %v")
	ErrLikeExists         = fmt.Errorf("like already exists")
	MsgArticleCreated     = "article successfully created"
	MsgArticleUpdated     = "article successfully updated"
	MsgArticleDeleted     = "article successfully deleted"
	MsgArticleLiked       = "article successfully liked"
	MsgArticleUnliked     = "article successfully unliked"

	// Validation messages
	ErrValidationFailed   = fmt.Errorf("validation failed: %v")
	ErrFieldRequired      = "field %s is required"
	ErrFieldMinLength     = "field %s must be at least %d characters"
	ErrFieldMaxLength     = "field %s must not exceed %d characters"
	ErrInvalidEmailFormat = "invalid email format"

	// Success messages
	MsgRegistrationSuccess = "registration successful"
	MsgLoginSuccess        = "login successful"

	// Server errors
	ErrInternalServer     = fmt.Errorf("internal server error")
	ErrDatabaseConnection = fmt.Errorf("database connection error")
	ErrDatabaseOperation  = fmt.Errorf("database operation failed")

	// Comment messages
	MsgCommentCreated  = "comment successfully created"
	MsgCommentUpdated  = "comment successfully updated"
	ErrGettingComments = fmt.Errorf("error to getting comments: %v")
	ErrCreatingComment = fmt.Errorf("error creating comment: %v")

	// User messages
	ErrGettingUser = fmt.Errorf("error getting user: %v")
)
