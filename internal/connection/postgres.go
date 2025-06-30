package connection

import (
	"fmt"
	"log"

	"my-golang-service-pos/internal/config"
	

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)



	func NewDb(cfg *config.Config, driver string) (*sqlx.DB, *gorm.DB, error) {
		if driver == "postgres" {
			return InitPostgres(cfg)
		} else if driver == "mysql" {
			// return InitMysql(cfg)
			log.Fatal("MySQL driver is not implemented yet")
		} else if driver == "sqlite" {
			// return InitSqlite(cfg)
			log.Fatal("SQLite driver is not implemented yet")
		} else if driver == "mongodb" {
			// return InitMongoDB(cfg)
			log.Fatal("MongoDB driver is not implemented yet")
		}
		return nil, nil, fmt.Errorf("unsupported driver: %s", driver)
	}

	func InitPostgres(cfg *config.Config) (db *sqlx.DB, dbGorm *gorm.DB, err error) {
		if cfg.PostgresDB == nil {
			log.Fatal("Postgres config is nil")
		}

		dbUser := cfg.PostgresDB.User
		dbPassword := cfg.PostgresDB.Password
		dbHost := cfg.PostgresDB.Host
		dbPort := cfg.PostgresDB.Port
		dbName := cfg.PostgresDB.DBName
		dbSSLMode := cfg.PostgresDB.SSLMode

		dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode)

		db = sqlx.MustConnect("postgres", dsn)

		if db == nil {
			log.Fatal("Failed to connect to PostgreSQL")
			panic("Error DB")
		}

		err = db.Ping()
		if err != nil {
			return
		}

		db.SetMaxOpenConns(100)
		db.SetMaxIdleConns(10)
		db.SetConnMaxLifetime(0) 

		dbGorm, err = gorm.Open(postgres.New(postgres.Config{
			Conn: db.DB,
		}), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: false,
			},
		},
		)


		if err != nil {
			log.Fatal("Failed to auto migrate PostgreSQL with GORM:", err)
		}

		// migration database

		err = dbGorm.AutoMigrate(
			// &model.User{},
			// &model.Product{},
			// &model.Transaction{},
			// &model.TransactionItems{},
			// &model.Cart{},
		)

		if err != nil {
			log.Fatal("Failed to auto migrate PostgreSQL with GORM:", err)
		}


	return db, dbGorm, nil

}
