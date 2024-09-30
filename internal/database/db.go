package db

import (
	"database/sql"
	"fmt"

	"github.com/XenHunt/go-test-project/internal/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type RTokens struct {
	bun.BaseModel `bun:"table:refresh_tokens,alias:rtokens"`
	Token         string
}

func MakeConection(dbConfig config.DataBaseConfig) *bun.DB {
	dsn := fmt.Sprintf("postgresql://%s:%s@localhost:%s/%s?sslmode=disabled", dbConfig.User, dbConfig.Password, dbConfig.Port, dbConfig.Database)
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())
	return db
}

// func createSchema(db *pg.DB) {

// }
