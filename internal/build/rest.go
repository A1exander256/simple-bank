package build

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/A1exander256/simple-bank/internal/restapi"
	openapi "github.com/A1exander256/simple-bank/internal/restapi/go"
	"github.com/A1exander256/simple-bank/internal/service/api/user"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

func (b *Builder) buildAPI() (*mux.Router, *loads.Document, error) {
	swaggerSpec, err := loads.Spec("api/rest/file.yaml")
	if err != nil {
		return nil, nil, fmt.Errorf("loading swagger specs: %w", err)
	}

	userService := user.NewService()

	h := restapi.NewHandler(userService)
	c := openapi.NewDefaultAPIController(h)
	router := openapi.NewRouter(c)

	return router, swaggerSpec, nil
}

func (b *Builder) RestAPIServer(ctx context.Context) (*http.Server, error) {
	router, swaggerSpec, err := b.buildAPI()
	if err != nil {
		return nil, fmt.Errorf("building API: %w", err)
	}

	server, err := b.HTTPServer(ctx, router)
	if err != nil {
		return nil, fmt.Errorf("crating http server: %w", err)
	}

	apiRouter := router.Name("api").Subrouter()

	//nolint:exhaustruct
	swaggerUIOpts := middleware.SwaggerUIOpts{
		BasePath: swaggerSpec.BasePath(),
		SpecURL:  fmt.Sprintf("%s/swagger.json", swaggerSpec.BasePath()),
	}

	apiRouter.PathPrefix(swaggerSpec.BasePath()).Handler(
		func() http.Handler {
			return middleware.Spec(
				swaggerSpec.BasePath(),
				swaggerSpec.Raw(),
				middleware.SwaggerUI(
					swaggerUIOpts,
					nil,
				),
			)
		}(),
	)

	return server, nil
}

func (b *Builder) HTTPServer(ctx context.Context, router http.Handler) (*http.Server, error) {
	const timeout = time.Millisecond * 25

	//nolint:exhaustruct
	server := http.Server{
		Addr:              b.config.HTTP.Addr(),
		ReadHeaderTimeout: timeout,
		Handler:           router,
		ErrorLog:          log.New(zerolog.Nop(), "", 0),
		BaseContext: func(net.Listener) context.Context {
			return ctx
		},
	}

	b.shutdown.add(server.Shutdown)

	return &server, nil
}
