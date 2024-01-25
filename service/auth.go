package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/deadlorpa/auth-app/interfaces"
	"github.com/deadlorpa/auth-app/model"
	"github.com/dgrijalva/jwt-go"
)

// TODO: вынести конфиг
const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}

type AuthService struct {
	Authorization interfaces.AuthorizationRepository
}

func NewAuthService(repository interfaces.AuthorizationRepository) *AuthService {
	return &AuthService{Authorization: repository}
}

func (s *AuthService) CreateUser(user model.User) (string, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.Authorization.CreateUser(user)
}

func (s *AuthService) GetToken(userSignIn model.UserSignIn) (string, error) {
	userSignIn.Password = generatePasswordHash(userSignIn.Password)
	user, err := s.Authorization.GetUser(userSignIn)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
