package db

import (
	"log"

	"github.com/XenHunt/go-test-project/internal/config"
	"github.com/go-pg/pg/v10"
)

type RTokens struct {
	Token string
}

func MakeConection(dbConfig config.DataBaseConfig) *pg.DB {
	opt, err := pg.ParseURL("postgres://" + dbConfig.User + ":" + dbConfig.Password + "@localhost:" + dbConfig.Port + "/" + dbConfig.Database)
	if err != nil {
		log.Fatal("Can not parse to PostgreSQL db")
	}
	db := pg.Connect(opt)

	return db
}

func createSchema(db *pg.DB) {

}
