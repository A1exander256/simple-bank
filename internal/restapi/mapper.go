package restapi

import (
	openapi "github.com/A1exander256/simple-bank/internal/restapi/go"
	"github.com/A1exander256/simple-bank/internal/service/api/user"
)

func mapUserPostFromRest(body openapi.UserPostRequest) *user.CreateUserParams {
	return &user.CreateUserParams{
		Username: body.Username,
		Password: body.Password,
		Email:    body.Email,
	}
}
