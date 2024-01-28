package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/deadlorpa/auth-app/appconfig"
	"github.com/deadlorpa/auth-app/interfaces"
	"github.com/deadlorpa/auth-app/model"
	"github.com/dgrijalva/jwt-go"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}

type AuthService struct {
	Authorization interfaces.AuthorizationRepository
}

func NewAuthService(repository interfaces.AuthorizationRepository) *AuthService {
	return &AuthService{
		Authorization: repository,
	}
}

func (s *AuthService) CreateUser(user model.User) (id string, err error) {
	config, err := appconfig.Get()
	if err != nil {
		return "", err
	}
	user.Password = generatePasswordHash(user.Password, config.AuthConfig.SHASalt)
	return s.Authorization.CreateUser(user)
}

func (s *AuthService) SignIn(userSignIn model.UserSignInRequest) (responce model.UserSignInResponce, err error) {
	var role model.Role

	config, err := appconfig.Get()
	if err != nil {
		return responce, err
	}

	userSignIn.Password = generatePasswordHash(userSignIn.Password, config.AuthConfig.SHASalt)
	user, err := s.Authorization.GetUser(userSignIn)
	if err != nil {
		return responce, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(config.AuthConfig.JWTTokenTTL) * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	role, err = s.Authorization.GetRole(user.IdRole)
	if err != nil {
		return responce, err
	}

	responce.Id = user.Id
	responce.Username = user.Username
	responce.Role = role
	responce.Token, err = token.SignedString([]byte(config.AuthConfig.JWTSigningKey))

	return responce, err
}

func generatePasswordHash(password string, salt string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
