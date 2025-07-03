package services

import (
	"context"
	"fmt"
	"restapp/config"
	"restapp/internal/messages"
	"restapp/internal/models"
	"restapp/internal/repositories"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceInterface interface {
	Register(ctx context.Context, req *models.RegisterRequest) (*models.UserResponse, error)
	Login(ctx context.Context, req *models.LoginRequest) (string, error)
	ValidateToken(tokenString string) (*models.Claims, error)
	FormatToken(tokenString string) string
}

type AuthService struct {
	r   repositories.AuthRepositoryInterface
	cfg *config.Config
}

func NewAuthService(r repositories.AuthRepositoryInterface, cfg *config.Config) *AuthService {
	return &AuthService{r: r, cfg: cfg}
}

func (s *AuthService) Register(ctx context.Context, user *models.RegisterRequest) (*models.UserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf(messages.ErrHashingPassword, err)
	}

	userModel := models.User{
		Username:  user.Username,
		Password:  string(hashedPassword),
		Email:     user.Email,
		Role:      "user",
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	registeredUser, err := s.r.Register(ctx, &userModel)
	if err != nil {
		return nil, fmt.Errorf(messages.ErrCreatingUser, err)
	}

	token, err := s.generateToken(registeredUser.Id)
	if err != nil {
		return nil, fmt.Errorf(messages.ErrGeneratingToken, err)
	}

	return &models.UserResponse{
		User:  registeredUser,
		Token: token,
	}, nil
}

func (s *AuthService) Login(ctx context.Context, user *models.LoginRequest) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	userModel, err := s.r.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return "", fmt.Errorf(messages.ErrGettingUser, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(user.Password))
	if err != nil {
		return "", fmt.Errorf(messages.ErrComparingPasswords, err)
	}

	return s.generateToken(userModel.Id)
}

func (s *AuthService) generateToken(userId int) (string, error) {
	expirationTime, err := strconv.Atoi(s.cfg.JWT.Expiration)
	if err != nil {
		return "", fmt.Errorf(messages.ErrConvertingExpTime, err)
	}

	claims := &models.Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expirationTime) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.cfg.JWT.Secret))
	if err != nil {
		return "", fmt.Errorf(messages.ErrSigningToken, err)
	}

	return tokenString, nil
}

func (s *AuthService) ValidateToken(tokenString string) (*models.Claims, error) {
	claims := &models.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(messages.ErrInvalidSigningMethod)
		}
		return []byte(s.cfg.JWT.Secret), nil
	})

	if err != nil {
		return nil, fmt.Errorf(messages.ErrParsingToken, err)
	}

	if !token.Valid {
		return nil, fmt.Errorf(messages.ErrInvalidToken)
	}

	return claims, nil
}

func (s *AuthService) FormatToken(tokenString string) string {
	const prefix = "Bearer "
	if len(tokenString) > len(prefix) && tokenString[:len(prefix)] == prefix {
		return tokenString[len(prefix):]
	}
	return ""
}
