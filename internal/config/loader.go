package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	expInt, _ := strconv.Atoi(os.Getenv("JWT_EXP"))
	readTimeout, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))
	writeTimeout, _ := strconv.Atoi(os.Getenv("SERVER_WRITE_TIMEOUT"))
	redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	appDebug := os.Getenv("APP_DEBUG") == "true"

	return &Config{
		App: App{
			Name:    os.Getenv("APP_NAME"),
			Env:     os.Getenv("APP_ENV"),
			Version: os.Getenv("APP_VERSION"),
			Debug:   appDebug,
		},
		Server: Server{
			Host:         os.Getenv("SERVER_HOST"),
			Port:         os.Getenv("SERVER_PORT"),
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
		},
		Jwt: Jwt{
			Key: os.Getenv("JWT_KEY"),
			Exp: expInt,
		},
		Log: Log{
			Level:    os.Getenv("LOG_LEVEL"),
			FilePath: os.Getenv("LOG_FILE_PATH"),
		},

		MongoDB: &MongoDB{
			URI:      os.Getenv("MONGO_URI"),
			Database: os.Getenv("MONGO_DATABASE"),
		},

		PostgresDB: &PostgresDB{
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
			DBName:   os.Getenv("POSTGRES_DB"),
			SSLMode: os.Getenv("POSTGRES_SSLMODE"),
		},

		Redis: &Redis{
			Host:     os.Getenv("REDIS_HOST"),
			Port:     os.Getenv("REDIS_PORT"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       redisDB,
		},
	}

}
