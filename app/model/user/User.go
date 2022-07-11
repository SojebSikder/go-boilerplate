package user

import (
	"context"
	"fmt"
	"log"

	"github.com/SojebSikder/goframe/ent"
)

// create user
func CreateUser(ctx context.Context, client *ent.Client) *ent.User {
	u, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		Save(ctx)
	if err != nil {
		return nil
	}
	log.Println("user was created: ", u)
	return nil
}

// get user

func QueryUser(ctx context.Context, client *ent.Client) (any, error) {
	u, err := client.User.
		Query().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}
