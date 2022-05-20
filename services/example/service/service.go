package service

import (
	"errors"

	"github.com/rs/zerolog"
	gw "github.com/timstudd/go-example/generated/protobuf/services/example"
	"github.com/timstudd/go-example/services/example/pkg/time"
	"google.golang.org/grpc"
)

// Config provides configuration for a service instance.
type Config struct {
	Logger     *zerolog.Logger
	TimeClient time.Client
}

// Service describes the interface for interacting with a service instance.
type Service interface {
	gw.ExampleServer
	LogsGRPCUnaryServerInterceptor() grpc.UnaryServerInterceptor
}

// New creates and returns a new service instance.
func New(conf *Config) (Service, error) {
	if conf.Logger == nil {
		return nil, errors.New("invalid logger")
	}
	if conf.TimeClient == nil {
		return nil, errors.New("invalid time client")
	}

	return &service{
		logger:     conf.Logger,
		timeClient: conf.TimeClient,
	}, nil
}

type service struct {
	gw.UnimplementedExampleServer

	logger     *zerolog.Logger
	timeClient time.Client
}
