package main

import (
	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

// containing all information need for ORM to connect to the PostgreSQL database
type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

// getPGOptions takes the given database options loaded from the .env file and returns an options struct for the ORM
func (db *Database) PGOptions() *pg.Options {
	return &pg.Options{
		Addr:     db.Host + ":" + db.Port,
		User:     db.User,
		Password: db.Password,
		Database: db.Database,
	}
}

type JWT struct {
	SecretKey               string
	HeaderField             string
	ValidityPeriodInMinutes int
}

// contains all settings of the application
type Config struct {
	DB   Database
	Port string
	Host string
	JWT  JWT
}

// readConfigFromEnviroment loads the config settings from the .env file
func readConfigFromEnviroment() Config {
	if err := godotenv.Load(".env"); err != nil {
		panic("Env file not found!")
	}

	jwtValidity, err := strconv.Atoi(os.Getenv("JWT_PERIOD_IN_MIN"))
	if err != nil {
		panic("Invalid JWT validity period! Not a valid integer!")
	}

	return Config{
		DB: Database{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_DATABASE"),
		},
		Port: os.Getenv("API_PORT"),
		Host: os.Getenv("API_HOST"),
		JWT: JWT{
			SecretKey:               os.Getenv("JWT_SECRET"),
			HeaderField:             os.Getenv("JWT_HEADER_FIELD"),
			ValidityPeriodInMinutes: jwtValidity,
		},
	}
}
