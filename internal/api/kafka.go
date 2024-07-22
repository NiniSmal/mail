package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"gitlab.com/nina8884807/mail/internal/service"
	gen "gitlab.com/nina8884807/mail/proto"
	"log/slog"
)

type KafkaHandler struct {
	service *service.SendService
	conn    *kafka.Reader
	l       *slog.Logger
}

func NewKafkaHandler(conn *kafka.Reader, service *service.SendService, l *slog.Logger) *KafkaHandler {
	return &KafkaHandler{
		service: service,
		conn:    conn,
		l:       l,
	}
}

func (k *KafkaHandler) OnCreateUser() {
	for {
		err := func() error {
			mail, err := k.conn.ReadMessage(context.Background())
			if err != nil {
				return fmt.Errorf("read message: %w", err)
			}

			var msg gen.SendEmailRequest

			err = json.Unmarshal(mail.Value, &msg)
			if err != nil {
				return fmt.Errorf("unmarshal: %w", err)
			}
			k.l.Info("read message", "kafka_msg", msg.String())
			err = k.service.SendMessage(&msg)
			if err != nil {
				return fmt.Errorf("send message: %w", err)
			}
			err = k.conn.CommitMessages(context.Background(), mail)
			if err != nil {
				return fmt.Errorf("commit messages: %w", err)
			}
			return nil

		}()
		if err != nil {
			k.l.Error("create user", "error", err)
			continue
		}
	}
}
