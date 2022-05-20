package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/timstudd/go-example/services/example/config"
	"github.com/timstudd/go-example/services/example/pkg/time/google"
	"github.com/timstudd/go-example/services/example/service"
	"github.com/timstudd/go-example/services/example/transport/grpc"
	"github.com/timstudd/go-example/services/example/transport/http"
)

func main() {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	var err error
	defer func() {
		if err != nil {
			logger.Fatal().Err(err).Msg("Error initializing server")
		}
	}()

	conf := config.Get()

	logger.Debug().
		Str("http_endpoint", conf.GoogleHTTPEndpoint).
		Msg("Initializing Google time client")
	timeClient, err := google.New(&google.Config{
		HTTPEndpoint: conf.GoogleHTTPEndpoint,
	})
	if err != nil {
		return
	}

	// logger.Debug().Msg("Initializing local time client")
	// timeClient, err := local.New(&local.Config{})
	// if err != nil {
	// 	return
	// }

	logger.Debug().Msg("Initializing service instance")
	server, err := service.New(&service.Config{
		Logger:     &logger,
		TimeClient: timeClient,
	})
	if err != nil {
		return
	}

	go func() {
		logger.Info().Str("listen_addr", conf.HTTPListenAddr).Msg("Creating HTTP listener")
		httpListener, err := http.New(&http.Config{
			GRPCListenPort:    conf.GRPCListenPort,
			HTTPListenAddress: conf.HTTPListenAddr,
			Logger:            &logger,
			Server:            server,
		})
		if err != nil {
			logger.Fatal().Err(err).Msg("Error creating HTTP listener")
		}
		logger.Fatal().Err(httpListener.Listen()).Msg("HTTP server died")
	}()

	go func() {
		logger.Info().Str("listen_addr", conf.GRPCListenAddr).Msg("Creating gRPC listener")
		grpcListener, err := grpc.New(&grpc.Config{
			GRPCListenAddress: conf.GRPCListenAddr,
			GRPCListenPort:    conf.GRPCListenPort,
			Logger:            &logger,
			LogsMiddleware:    server.LogsGRPCUnaryServerInterceptor(),
			Server:            server,
		})
		if err != nil {
			logger.Fatal().Err(err).Msg("Error creating gRPC listener")
		}
		logger.Fatal().Err(grpcListener.Listen()).Msg("gRPC server died")
	}()

	select {}
}
