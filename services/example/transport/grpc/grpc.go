package grpc

import (
	"errors"
	"math"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/rs/zerolog"
	gw "github.com/timstudd/go-example/generated/protobuf/services/example"
	"google.golang.org/grpc"
)

// Config represents the configuration for a listener.
type Config struct {
	GRPCListenAddress string
	GRPCListenPort    uint
	Logger            *zerolog.Logger
	LogsMiddleware    grpc.UnaryServerInterceptor
	Server            gw.ExampleServer
}

// Listener describes the interface for interacting with a listener.
type Listener interface {
	Listen() error
}

// New creates and returns a new listener.
func New(conf *Config) (Listener, error) {
	if conf.GRPCListenAddress == "" {
		return nil, errors.New("invalid gRPC listen address")
	}
	if conf.GRPCListenPort == 0 {
		return nil, errors.New("invalid gRPC listen port")
	}
	if conf.Logger == nil {
		return nil, errors.New("invalid logger")
	}
	if conf.LogsMiddleware == nil {
		return nil, errors.New("invalid logs middleware")
	}
	if conf.Server == nil {
		return nil, errors.New("invalid server")
	}

	return &listener{
		ExampleServer: conf.Server,

		grpcListenAddress: conf.GRPCListenAddress,
		grpcListenPort:    conf.GRPCListenPort,
		logger:            conf.Logger,
		logsMiddleware:    conf.LogsMiddleware,
	}, nil
}

type listener struct {
	gw.ExampleServer

	grpcListenAddress string
	grpcListenPort    uint
	logger            *zerolog.Logger
	logsMiddleware    grpc.UnaryServerInterceptor
}

func (l *listener) Listen() error {
	listener, err := net.Listen("tcp", l.grpcListenAddress)
	if err != nil {
		return err
	}

	middleware := grpc_middleware.ChainUnaryServer(
		l.logsMiddleware,
	)

	grpcServerOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(middleware),
		grpc.MaxConcurrentStreams(math.MaxUint32),
		grpc.MaxRecvMsgSize(math.MaxUint32),
		grpc.MaxSendMsgSize(math.MaxUint32),
	}

	grpcServe := grpc.NewServer(grpcServerOptions...)

	gw.RegisterExampleServer(grpcServe, l)

	return grpcServe.Serve(listener)
}
