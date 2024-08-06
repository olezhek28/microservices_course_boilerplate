package usecases

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/IBM/sarama"
	"github.com/neracastle/go-libs/pkg/sys/logger"
	"golang.org/x/exp/slog"

	domain "github.com/neracastle/auth/internal/domain/user"
	def "github.com/neracastle/auth/internal/usecases/models"
)

// Create создает нового пользователя
func (s *Service) Create(ctx context.Context, req def.CreateDTO) (int64, error) {
	log := logger.GetLogger(ctx)
	log.Debug("called", slog.String("method", "usecases.Create"))

	var err error
	var newUser *domain.User

	if req.Password != req.PasswordConfirm {
		return 0, errors.New("пароли не совпадают")
	}

	if req.IsAdmin {
		newUser, err = domain.NewAdmin(req.Email, req.Password, req.Name)
	} else {
		newUser, err = domain.NewUser(req.Email, req.Password, req.Name)
	}

	if err != nil {
		return 0, err
	}

	err = s.usersRepo.Save(ctx, newUser)
	if err != nil {
		log.Error("failed to create user", slog.String("error", err.Error()), slog.String("method", "usecases.Create"))
		return 0, err
	}

	jsonStr, err := json.Marshal(newUser)
	if err != nil {
		log.Error("failed to marshal user", slog.String("error", err.Error()), slog.String("method", "usecases.Create"))
		return 0, err
	}

	partition, offset, err := s.producer.SendMessage(&sarama.ProducerMessage{
		Topic: s.Config.NewUserTopic,
		Value: sarama.ByteEncoder(jsonStr),
	})
	if err != nil {
		log.Error("failed to send message to kafka", slog.String("error", err.Error()), slog.String("method", "usecases.Create"))
		return 0, err
	}

	log.Debug("message sent to kafka", slog.Int("partition", int(partition)), slog.Int64("offset", offset))

	return newUser.ID, nil
}
