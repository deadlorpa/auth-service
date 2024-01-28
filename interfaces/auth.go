package interfaces

import "github.com/deadlorpa/auth-app/model"

type AuthorizationService interface {
	SignUp(user model.User) (id string, err error)
	SignIn(userSignIn model.UserSignInRequest) (responce model.UserSignInResponce, err error)
}
