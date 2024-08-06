package kafka

import (
	"context"
	"errors"

	"github.com/IBM/sarama"
)

var _ sarama.ConsumerGroupHandler = (*GroupHandler)(nil)

// MessageHandler обработчик сообщения
type MessageHandler func(ctx context.Context, msg *sarama.ConsumerMessage) error

// LifecycleHandler обработчик событий жизненного цикла сессии
type LifecycleHandler func(sarama.ConsumerGroupSession) error

// ErrRebalancingGroup возникает когда в группу входит consumer и возникает перебалансировка группы
var ErrRebalancingGroup = errors.New("rebalancing group started")

// ErrHandlerError возникает, когда пользовательский обработчик вернул ошибку
var ErrHandlerError = errors.New("handler returned error")

// GroupHandler обработчик
type GroupHandler struct {
	setupHandler   LifecycleHandler
	cleanupHandler LifecycleHandler
	msgHandler     MessageHandler
}

// NewGroupHandler новый экзмепляр
func NewGroupHandler() *GroupHandler {
	return &GroupHandler{}
}

// SetMessageHandler задает пользовательский метод для обработки сообщения
func (h *GroupHandler) SetMessageHandler(handler MessageHandler) {
	h.msgHandler = handler
}

// Setup запускается в начале новой сессии (однократно, при присоединении к группе)
func (h *GroupHandler) Setup(session sarama.ConsumerGroupSession) error {
	if h.setupHandler != nil {
		return h.setupHandler(session)
	}

	return nil
}

// Cleanup запускается в конце жизни сессии после того как все горутины ConsumeClaim завершаться
func (h *GroupHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	if h.cleanupHandler != nil {
		return h.cleanupHandler(session)
	}

	return nil
}

// ConsumeClaim обрабатывает поступающие сообщения (вызов блокирующий, sarama вызовет этот метод в горутине)
func (h *GroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				return ErrRebalancingGroup
			}

			if h.msgHandler != nil {
				err := h.msgHandler(session.Context(), message)
				if err != nil {
					return errors.Join(ErrHandlerError, err)
				}
			}

			session.MarkMessage(message, "")

		case <-session.Context().Done():
			return ErrRebalancingGroup
		}
	}
}
