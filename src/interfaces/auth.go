package interfaces

import (
	"github.com/hendralatumeten/vehicles_rental/src/database/orm/models"
	"github.com/hendralatumeten/vehicles_rental/src/libs"
)

type AuthService interface {
	Login(body models.User) *libs.Response
	Register(body models.User) *libs.Response
}
