package usecases

import (
	"context"
	"time"

	"github.com/IBM/sarama"
	"github.com/neracastle/go-libs/pkg/db"

	"github.com/neracastle/auth/internal/kafka"
	"github.com/neracastle/auth/internal/repository/action"
	"github.com/neracastle/auth/internal/repository/user"
	def "github.com/neracastle/auth/internal/usecases/models"
)

// UserService возможные сценарии с пользователем
type UserService interface {
	Create(ctx context.Context, req def.CreateDTO) (int64, error)
	Update(ctx context.Context, user def.UpdateDTO) error
	Get(ctx context.Context, userID int64) (def.UserDTO, error)
	Delete(ctx context.Context, userID int64) error
}

// Service сервис сценарием пользователя
type Service struct {
	usersRepo   user.Repository
	usersCache  user.Cache
	actionsRepo action.Repository
	db          db.DB
	producer    sarama.SyncProducer
	consumer    kafka.Consumer
	Config
}

// Config параметры сервиса
type Config struct {
	// время кэширования данных о пользователе
	CacheTTL time.Duration
	// топик для отправки событий пользователя
	NewUserTopic string
}

// NewService новый экзмепляр usecase-сервиса
func NewService(usersRepo user.Repository,
	usersCache user.Cache,
	actionsRepo action.Repository,
	db db.DB,
	producer sarama.SyncProducer,
	consumer kafka.Consumer,
	config Config) *Service {
	return &Service{
		usersRepo:   usersRepo,
		usersCache:  usersCache,
		actionsRepo: actionsRepo,
		db:          db,
		producer:    producer,
		consumer:    consumer,
		Config: Config{
			CacheTTL:     config.CacheTTL,
			NewUserTopic: config.NewUserTopic,
		},
	}
}
