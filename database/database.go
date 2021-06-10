package database

import (
	"database/sql"
	"os"

	"entgo.io/ent/dialect"
	"github.com/y-yagi/go-api-template/ent"

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
	dbgDrv := dialect.Debug(drv)
	Client = ent.NewClient(ent.Driver(dbgDrv))
	return nil
}
