package postgresdb

import (
	"abdullayev13/timeup/internal/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func New() *gorm.DB {
	dsn := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%d sslmode=disable`,
		config.DbHost, config.DbUsername, config.DbPassword, config.DbName, config.DbPort)
	if config.IsDbBlank() {
		dsn = "postgres://tfpjcgxs:044_2u9xKgU0qxU0etYHKFoMT5_SQ_4s@john.db.elephantsql.com/tfpjcgxs"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
