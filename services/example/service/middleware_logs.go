package service

import (
	"context"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// LogsGRPCUnaryServerInterceptor intercepts exceptions and returns standard response errors.
func (s *service) LogsGRPCUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		res, err := handler(ctx, req)
		if err != nil {
			grpcCode := status.Code(err)
			if grpcCode == codes.Unknown || grpcCode == codes.Internal {
				s.logger.Error().Err(err).Msg("Unhandled error returned")
				err = status.Error(codes.Internal, http.StatusText(http.StatusInternalServerError))
			}

			return nil, err
		}

		return res, nil
	}
}
