package config

import (
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)




type (

	Config struct {
		App        App
		Server     Server
		Jwt        Jwt
		Log        Log
		MongoDB    *MongoDB    // Optional
		PostgresDB *PostgresDB // Optional
		Redis      *Redis      // Optional
	}
	
	App struct {
		Name    string
		Env     string // development | production | testing
		Version string
		Debug   bool
	}
	
	Server struct {
		Host         string
		Port         string
		ReadTimeout  int // in seconds
		WriteTimeout int
	}
	
	Log struct {
		Level    string // debug | info | warn | error
		FilePath string
	}
	
	MongoDB struct {
		URI      string
		Database string
	}
	
	PostgresDB struct {
		User     string
		Password string
		Host     string
		Port     string
		DBName   string
		SSLMode  string
	}
	
	Redis struct {
		Host     string
		Port     string
		Password string
		DB       int
	}
	
	Jwt struct {
		Key string
		Exp int
	}
	
	
	ServiceParam struct {
		RepoParam
		Db *sqlx.DB
		DbGorm *gorm.DB
	}
	
	
	RepoParam struct {
		Config *Config
	
	}
)