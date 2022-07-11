package orm

import (
	"context"

	"github.com/SojebSikder/goframe/ent"
)

// Ctx is the context of the ORM
var Ctx context.Context

// Client is the client of the ORM
var Client *ent.Client

// Init initializes the ORM
func Init(ctx context.Context, client *ent.Client) {
	Ctx = ctx
	Client = client
}
