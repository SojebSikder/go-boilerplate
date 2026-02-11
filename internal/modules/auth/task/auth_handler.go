package authtask

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

type AuthProcessHandler struct {
	logger *zap.Logger
}

func NewAuthProcessHandler(
	logger *zap.Logger,
) *AuthProcessHandler {
	return &AuthProcessHandler{
		logger: logger,
	}
}

func (s *AuthProcessHandler) Handle(ctx context.Context, t *asynq.Task) error {
	s.logger.Info("[Asynq] Processing auth task", zap.String("payload", string(t.Payload())))

	//
	// prepare payload
	//
	var payload AuthPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		s.logger.Error("Failed to unmarshal payload", zap.Error(err))
		return err
	}

	fmt.Println("Hello, " + payload.Name)

	return nil
}
