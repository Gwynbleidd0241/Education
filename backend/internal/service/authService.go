package service

import (
	"Kursash/internal/models"
	"Kursash/internal/repository"
	"crypto/sha1"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const (
	salt       = "hjq123hjqw124117ajf4ajs"
	signingKey = "qrkjk123#4#%35FSFJlja#4353123KSFjH"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.UserModel) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
