package api

import (
	"context"
	"fmt"
	gen "read-massage/proto"
)

type Handler struct {
	*gen.UnimplementedMailServer
}

func (h *Handler) SendEmail(ctx context.Context, request *gen.SendEmailRequest) (*gen.SendEmailResponse, error) {
	fmt.Println("got request for ", request.To)
	return &gen.SendEmailResponse{}, nil
}
