package postgresdb

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func New() *gorm.DB {
	dsn := "host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable"
	// TODO make real postgres db
	dsn = "postgres://tfpjcgxs:044_2u9xKgU0qxU0etYHKFoMT5_SQ_4s@john.db.elephantsql.com/tfpjcgxs"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
