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

	if os.Getenv("MYSQL_HOST") == "" {
		fmt.Printf("Environment variable MYSQL_HOST is not set")
	}
	if os.Getenv("MYSQL_USER") == "" {
		fmt.Printf("Environment variable MYSQL_USER is not set")
	}
	if os.Getenv("MYSQL_PORT") == "" {
		fmt.Printf("Environment variable MYSQL_PORT is not set")
	}
	if os.Getenv("MYSQL_DATABASE") == "" {
		fmt.Printf("Environment variable MYSQL_DATABASE is not set")
	}

	dbHost := os.Getenv("MYSQL_HOST")
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbPort := os.Getenv("MYSQL_PORT")
	dbDatabase := os.Getenv("MYSQL_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbDatabase)

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
