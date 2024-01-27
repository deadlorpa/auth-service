package interfaces

import "github.com/deadlorpa/auth-app/model"

// TODO: Разнести репозитории по сущностям

type AuthorizationRepository interface {
	CreateUser(user model.User) (id string, err error)
	GetUser(userSignIn model.UserSignInRequest) (user model.User, err error)
	GetRole(id string) (role model.Role, err error)
}

type AuthorizationService interface {
	CreateUser(user model.User) (id string, err error)
	SignIn(userSignIn model.UserSignInRequest) (responce model.UserSignInResponce, err error)
}
