package service

import (
	"context"

	gw "github.com/timstudd/go-example/generated/protobuf/services/example"
)

func (s *service) Echo(ctx context.Context, req *gw.EchoRequest) (*gw.EchoResponse, error) {
	currTime, err := s.timeClient.GetCurrent()
	if err != nil {
		return nil, err
	}

	return &gw.EchoResponse{
		CurrentTime: uint32(currTime.Unix()),
		Message:     req.Message,
	}, nil
}
