package config

import "fmt"

type HTTP struct {
	Port int32 `envconfig:"HTTP_PORT" default:"8080"`
}

func (h *HTTP) Addr() string {
	return fmt.Sprintf(":%d", h.Port)
}
