package http

import (
	"context"
	"errors"
	"fmt"
	"math"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog"
	gw "github.com/timstudd/go-example/generated/protobuf/services/example"
	"github.com/timstudd/go-example/services/example/service"
	"google.golang.org/grpc"
)

// Config represents the configuration for a listener.
type Config struct {
	GRPCListenPort    uint
	HTTPListenAddress string
	Logger            *zerolog.Logger
	Server            service.Service
}

// Listener describes the interface for interacting with a listener.
type Listener interface {
	Listen() error
}

// New creates and returns a new listener.
func New(conf *Config) (Listener, error) {
	if conf.GRPCListenPort == 0 {
		return nil, errors.New("invalid gRPC listen port")
	}
	if conf.HTTPListenAddress == "" {
		return nil, errors.New("invalid HTTP listen address")
	}
	if conf.Logger == nil {
		return nil, errors.New("invalid logger")
	}
	if conf.Server == nil {
		return nil, errors.New("invalid server")
	}

	return &listener{
		grpcListenPort:    conf.GRPCListenPort,
		httpListenAddress: conf.HTTPListenAddress,
		logger:            conf.Logger,
		server:            conf.Server,
	}, nil
}

type listener struct {
	grpcCompressMux   http.Handler
	grpcListenPort    uint
	grpcMux           *runtime.ServeMux
	httpListenAddress string
	logger            *zerolog.Logger
	server            service.Service
}

func (l *listener) Listen() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()

	mux.HandleFunc("/", l.handlerRoot)

	l.grpcMux = runtime.NewServeMux()

	l.grpcCompressMux = handlers.CompressHandler(l.grpcMux)

	localGRPCAddress := fmt.Sprintf("127.0.0.1:%d", l.grpcListenPort)

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxUint32),
			grpc.MaxCallSendMsgSize(math.MaxUint32),
		),
	}
	err := gw.RegisterExampleHandlerFromEndpoint(ctx, l.grpcMux, localGRPCAddress, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(l.httpListenAddress, mux)
}

func (l *listener) handlerRoot(w http.ResponseWriter, r *http.Request) {
	// Health check endpoint.
	if r.URL.Path == "/" {
		w.WriteHeader(http.StatusOK)
		return
	}

	l.grpcCompressMux.ServeHTTP(w, r)
}
