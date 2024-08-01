package restapi

import (
	"context"
	"net/http"

	openapi "github.com/A1exander256/simple-bank/internal/restapi/go"
	"github.com/A1exander256/simple-bank/internal/service/api/user"
)

type service struct {
	user *user.Service
}

type Handler struct {
	openapi.DefaultAPIServicer

	s *service
}

func NewHandler(userSrv *user.Service) *Handler {
	return &Handler{
		DefaultAPIServicer: openapi.NewDefaultAPIService(),
		s:                  &service{user: userSrv},
	}
}

func (h *Handler) UserPost(ctx context.Context, body openapi.UserPostRequest) (openapi.ImplResponse, error) {
	guid, err := h.s.user.CreateUser(ctx, mapUserPostFromRest(body))
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, openapi.Error{
			Message: "server error",
		}), err
	}

	return openapi.Response(http.StatusCreated, openapi.UserPost201Response{
		Guid: guid.String(),
	}), nil
}
