package main

import (
	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
	"os"
)

// containing all information need for ORM to connect to the PostgreSQL database
type Database struct {
	Host string
	Port string
	User string
	Password string
	Database string
}

// getPGOptions takes the given database options loaded from the .env file and returns an options struct for the ORM
func (db *Database) PGOptions() *pg.Options {
	return &pg.Options{
		Addr: db.Host + ":" + db.Port,
		User: db.User,
		Password: db.Password,
		Database: db.Database,
	}
}

// contains all settings of the application
type Config struct {
	DB Database
	Port string
	Host string
}

// readConfigFromEnviroment loads the config settings from the .env file
func readConfigFromEnviroment() Config {
	if err := godotenv.Load(".env"); err != nil {
		panic("Env file not found!")
	}

	return Config{
		DB:   Database{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_DATABASE"),
		},
		Port: os.Getenv("API_PORT"),
		Host: os.Getenv("API_HOST"),
	}
}