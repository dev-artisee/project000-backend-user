package datasource

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"project000-backend-user/config"
)

func NewGORMDatabase(cfg config.Config) (*gorm.DB, error) {
	dbConfig := cfg.Database

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
		dbConfig.Port,
		dbConfig.TimeZone,
	)

	gormConfig := gorm.Config{}

	// local 환경에서 SQL 로그를 출력하도록 설정
	if cfg.Server.Env == "local" {
		gormConfig.Logger = logger.Recorder.LogMode(logger.Info)
	}

	return gorm.Open(postgres.Open(dsn), &gormConfig)
}
