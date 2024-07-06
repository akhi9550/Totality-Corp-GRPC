package db

import (
	"fmt"
	"grpc-user-service/pkg/config"
	"grpc-user-service/pkg/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=postgres port=%s password=%s sslmode=disable", cfg.DBHost, cfg.DBUser, cfg.DBPort, cfg.DBPassword)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	rows, err := db.Raw("SELECT 1 FROM pg_database WHERE datname = ?", cfg.DBName).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		query := fmt.Sprintf("CREATE DATABASE %s", cfg.DBName)
		if err := db.Exec(query).Error; err != nil {
			return nil, err
		}
		fmt.Printf("Database %s created successfully.\n", cfg.DBName)
	}
	psqlInfo = fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s sslmode=disable", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(
		&domain.User{},
	)
	return db, nil
}
