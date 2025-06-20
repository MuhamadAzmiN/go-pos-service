package connection

import (
	"fmt"
	"log"
	"my-golang-service-pos/internal/config"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func InitPostgres(cfg *config.Config) *gorm.DB {
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


	fmt.Println("POSTGRES INFO:", cfg.PostgresDB.Host)
	db , err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Postgres ping failed:", err)
	}

	log.Println("✅ Connected to PostgreSQL")

	return db
}
