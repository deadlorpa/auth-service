package interfaces

import "github.com/deadlorpa/auth-app/model"

type UsersRepository interface {
	Create(user model.User) (id string, err error)
	GetBySignIn(userSignIn model.UserSignInRequest) (user model.User, err error)
}
