package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/muhangga/internal/entity"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenConn() (*gorm.DB, error) {

	dbHost := os.Getenv("MYSQL_HOST")
	dbUser := os.Getenv("MYSQL_USER")
	dbPort := os.Getenv("MYSQL_PORT")
	dbDatabase := os.Getenv("MYSQL_DATABASE")

	dsn := fmt.Sprintf("%s:@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbHost, dbPort, dbDatabase)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&entity.Todo{}, &entity.Activities{})

	pool, err := db.DB()
	if err != nil {
		panic("Failed to get database connection pool")
	}

	pool.SetMaxIdleConns(10)
	pool.SetMaxOpenConns(100)
	pool.SetConnMaxLifetime(time.Hour)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := pool.PingContext(ctx); err != nil {
		log.Error().Msgf("Failed to ping database: %v", err)
	}

	fmt.Println("Database connected")

	return db, nil
}

func CloseConn(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Error().Msgf("cant close database: %v", err)
	}
	sqlDB.Close()
}
