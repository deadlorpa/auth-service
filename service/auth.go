package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/deadlorpa/auth-app/configs"
	"github.com/deadlorpa/auth-app/interfaces"
	"github.com/deadlorpa/auth-app/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}

type AuthService struct {
	config        configs.AuthConfig
	Authorization interfaces.AuthorizationRepository
}

func NewAuthService(repository interfaces.AuthorizationRepository) *AuthService {
	return &AuthService{
		config: configs.AuthConfig{
			SHASalt:       viper.GetString("auth.sha_salt"),
			JWTSigningKey: viper.GetString("auth.jwt_signing_key"),
			JWTTokenTTL:   viper.GetInt("auth.jwt_token_ttl"),
		},
		Authorization: repository,
	}
}

func (s *AuthService) CreateUser(user model.User) (string, error) {
	user.Password = generatePasswordHash(user.Password, s.config.SHASalt)
	return s.Authorization.CreateUser(user)
}

func (s *AuthService) GetToken(userSignIn model.UserSignIn) (string, error) {
	userSignIn.Password = generatePasswordHash(userSignIn.Password, s.config.SHASalt)
	user, err := s.Authorization.GetUser(userSignIn)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(s.config.JWTTokenTTL) * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(s.config.JWTSigningKey))
}

func generatePasswordHash(password string, salt string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
