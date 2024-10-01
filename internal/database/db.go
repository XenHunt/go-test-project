package db_module

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/XenHunt/go-test-project/internal/config"
	"golang.org/x/crypto/bcrypt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type RefreshToken struct {
	bun.BaseModel `bun:"table:refresh_tokens"`
	token         string
}

func MakeConection(dbConfig config.DataBaseConfig) *bun.DB {
	dsn := fmt.Sprintf("postgresql://%s:%s@localhost:%s/%s?sslmode=disabled", dbConfig.User, dbConfig.Password, dbConfig.Port, dbConfig.Database)
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())
	return db
}

func CreateSchema(db *bun.DB, ctx context.Context) error {
	_, err := db.NewCreateTable().Model((*RefreshToken)(nil)).Exec(ctx)
	return err
}

func AddToken(db *bun.DB, token string, ctx context.Context) error {
	hashedToken, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = db.NewInsert().Model(RefreshToken{token: string(hashedToken)}).Exec(ctx)
	return err
}

func DropToken(db *bun.DB, token string, ctx context.Context) error {
	hashedToken, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = db.NewDelete().Model((*RefreshToken)(nil)).Where("toke = ?", hashedToken).Exec(ctx)
	return err
}

func TokenExists(db *bun.DB, token string, ctx context.Context) bool {
	hashedToken, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return false
	}

	exists, err := db.NewSelect().Model((*RefreshToken)(nil)).Where("token = ?", hashedToken).Exists(ctx)

	return exists
}
