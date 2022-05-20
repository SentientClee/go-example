package config

import (
	"fmt"

	"github.com/timstudd/go-example/lib/utils/environment"
)

type Configuration struct {
	GRPCListenAddr string
	GRPCListenHost string
	GRPCListenPort uint

	GoogleHTTPEndpoint string

	HTTPListenAddr string
}

func Get() *Configuration {
	c := &Configuration{
		GRPCListenHost: environment.GetOrPanic("GRPC_LISTEN_HOST"),
		GRPCListenPort: uint(environment.GetIntOrPanic("GRPC_LISTEN_PORT")),

		GoogleHTTPEndpoint: environment.GetOrDefault("GOOGLE_HTTP_ENDPOINT", "https://www.google.com"),

		HTTPListenAddr: environment.GetOrPanic("HTTP_LISTEN_ADDR"),
	}

	c.GRPCListenAddr = fmt.Sprintf("%s:%d", c.GRPCListenHost, c.GRPCListenPort)

	return c
}
