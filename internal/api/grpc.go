package api

import (
	"context"
	"errors"
	"fmt"
	gen "gitlab.com/nina8884807/mail/proto"
	"gitlab.com/nina8884807/mail/service"
	"log"
)

type GrpcHandler struct {
	*gen.UnimplementedMailServer
	service *service.SendService
}

func NewGrpcHandler(s *service.SendService) *GrpcHandler {
	return &GrpcHandler{
		UnimplementedMailServer: &gen.UnimplementedMailServer{},
		service:                 s,
	}
}

func (h *GrpcHandler) SendEmail(ctx context.Context, request *gen.SendEmailRequest) (*gen.SendEmailResponse, error) {
	if request == nil {
		return nil, errors.New("request is empty")
	}
	err := h.service.SendMessage(request)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("got request for ", request.To, request.Text)
	return &gen.SendEmailResponse{}, nil

}
