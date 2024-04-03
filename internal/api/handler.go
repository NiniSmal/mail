package api

import (
	"context"
	"errors"
	"fmt"
	new_service "gitlab.com/nina8884807/mail/new-service"
	gen "gitlab.com/nina8884807/mail/proto"
	"log"
)

type Handler struct {
	*gen.UnimplementedMailServer
	service new_service.SendService
}

func (h *Handler) SendEmail(ctx context.Context, request *gen.SendEmailRequest) (*gen.SendEmailResponse, error) {
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
