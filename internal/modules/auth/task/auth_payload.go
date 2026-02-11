package authtask

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

const (
	TypeAuthProcess = "auth:process"
)

type AuthPayload struct {
	Name string `json:"name"`
}

func NewAuthTask(name string) (*asynq.Task, error) {
	payload := AuthPayload{
		Name: name,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(
		TypeAuthProcess,
		payloadBytes,
	), nil
}
