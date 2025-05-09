package messages

const (
	// Auth messages
	ErrInvalidCredentials    = "invalid email or password"
	ErrEmailAlreadyExists    = "email already exists"
	ErrUsernameAlreadyExists = "username already exists"
	ErrInvalidToken          = "invalid token"
	ErrMissingToken          = "missing authorization token"
	ErrParsingToken          = "error parsing token: %v"
	ErrTokenExpired          = "token expired"
	ErrInvalidSigningMethod  = "invalid signing method"
	ErrSigningToken          = "error signing token: %v"
	ErrHashingPassword       = "error hashing password: %v"
	ErrBadRequest            = "bad request"
	ErrConvertingExpTime     = "error converting expiration time: %v"
	ErrComparingPasswords    = "error comparing passwords: %v"

	// Article messages
	ErrFetchArticles      = "Failed to fetch articles: %s"
	ErrArticleNotFound    = "article not found"
	ErrInvalidArticleData = "invalid article data"
	ErrInvalidArticleID   = "invalid article ID"
	ErrGettingArticles    = "error getting articles: %v"
	MsgArticleCreated     = "article successfully created"
	MsgArticleUpdated     = "article successfully updated"
	MsgArticleDeleted     = "article successfully deleted"

	// Validation messages
	ErrValidationFailed   = "validation failed: %v"
	ErrFieldRequired      = "field %s is required"
	ErrFieldMinLength     = "field %s must be at least %d characters"
	ErrFieldMaxLength     = "field %s must not exceed %d characters"
	ErrInvalidEmailFormat = "invalid email format"

	// Success messages
	MsgRegistrationSuccess = "registration successful"
	MsgLoginSuccess        = "login successful"

	// Server errors
	ErrInternalServer     = "internal server error"
	ErrDatabaseConnection = "database connection error"
	ErrDatabaseOperation  = "database operation failed"

	// Comment messages
	MsgCommentCreated = "comment successfully created"
	MsgCommentUpdated = "comment successfully updated"

	// User messages
	ErrGettingUser = "error getting user: %v"
)
