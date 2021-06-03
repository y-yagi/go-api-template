package database

import (
	"context"
	"fmt"
	"os"

	"github.com/y-yagi/go-api-template/ent"
	"github.com/y-yagi/go-api-template/ent/migrate"

	_ "github.com/lib/pq"
)

var (
	Client *ent.Client
)

func New() error {
	var err error

	Client, err = ent.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return fmt.Errorf("failed connecting to DB: %v", err)
	}
	ctx := context.Background()

	// Run migration.
	err = Client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		return fmt.Errorf("failed creating schema resources: %v", err)
	}

	return nil
}
