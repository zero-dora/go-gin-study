package auth_service

import "github.com/zero-dora/go-gin-example/models"

type Auth struct {
	Username string
	Password string
}

func (a *Auth) Check() (bool, error) {
	return models.ChechAuth(a.Username, a.Password)
}
