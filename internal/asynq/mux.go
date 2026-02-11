package asynqfx

import (
	"github.com/hibiken/asynq"
	authtask "github.com/sojebsikder/go-boilerplate/internal/modules/auth/task"
)

func NewMux(
	authProcessHander *authtask.AuthProcessHandler,
) *asynq.ServeMux {
	mux := asynq.NewServeMux()
	// add worker handler here
	mux.HandleFunc(authtask.TypeAuthProcess, authProcessHander.Handle)

	return mux
}
