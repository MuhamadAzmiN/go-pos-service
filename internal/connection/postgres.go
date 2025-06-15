package connection

import (
	"database/sql"
	"fmt"
	"log"
	"my-echo-chat_service/internal/config"

	_ "github.com/lib/pq"
)

var PostgresDB *sql.DB

func InitPostgres(cfg *config.Config) *sql.DB {
	if cfg.PostgresDB == nil {
		log.Fatal("Postgres config is nil")
	}

	psqlInfo := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		cfg.PostgresDB.User,
		cfg.PostgresDB.Password,
		cfg.PostgresDB.DBName,
		cfg.PostgresDB.Host,
		cfg.PostgresDB.Port,
		cfg.PostgresDB.SSLMode,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("❌ Failed to open Postgres connection:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("❌ Postgres ping failed:", err)
	}

	log.Println("✅ Connected to PostgreSQL")
	PostgresDB = db

	return db
}
