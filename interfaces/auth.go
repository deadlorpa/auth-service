package interfaces

import "github.com/deadlorpa/auth-app/model"

type AuthorizationRepository interface {
	CreateUser(user model.User) (string, error)
	GetUser(userSignIn model.UserSignIn) (model.User, error)
}

type AuthorizationService interface {
	CreateUser(user model.User) (string, error)
	GetToken(userSignIn model.UserSignIn) (string, error)
}
