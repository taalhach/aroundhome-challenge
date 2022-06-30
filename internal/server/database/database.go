package database

import (
	"fmt"

	"github.com/taalhach/aroundhome-challennge/internal/server/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbSession struct {
	*gorm.DB
}

var Db DbSession

func MustConnectDB(cfg *configs.DatabaseConfig) {
	db, err := gorm.Open(postgres.Open(cfg.ConnString()), &gorm.Config{})
	if err != nil {
		fmt.Printf("Got error when connect database, the error is '%v'", err)
		panic(err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		fmt.Printf("db connection configs failed, the error is '%v'", err)
		panic(err)
	}

	sqlDb.SetMaxOpenConns(10)
	sqlDb.SetMaxIdleConns(5)

	// start debug mod
	if cfg.ShowSql {
		db.Debug()
	}

	db.Logger = logger.Default.LogMode(logger.Info)
	Db = DbSession{
		db,
	}
}

//Migrate run migration and takes models as argument
// gorm's auto migrator does not alter columns in some cases
func Migrate(models ...interface{}) error {
	return Db.AutoMigrate(models...)
}

