package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/Surbhibagad/cloud-native-ai-microservice-platform/services/auth-service/internal/config"
)

func Connect(cfg *config.Config) *pgx.Conn {

	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBSSLMode,
	)

	conn, err := pgx.Connect(context.Background(), connString)

	if err != nil {
		log.Fatal("Database Connection Failed:", err)
	}

	log.Println(" Connected to PostgreSQL")

	return conn
}