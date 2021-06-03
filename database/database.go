package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"entgo.io/ent/dialect"
	"github.com/y-yagi/go-api-template/ent"
	"github.com/y-yagi/go-api-template/ent/migrate"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var (
	Client *ent.Client
)

func New() error {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}

	drv := entsql.OpenDB(dialect.Postgres, db)
	Client = ent.NewClient(ent.Driver(drv))

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
