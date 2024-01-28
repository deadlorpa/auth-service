package interfaces

import "github.com/deadlorpa/auth-app/model"

type RolesRepository interface {
	GetById(id string) (role model.Role, err error)
}
