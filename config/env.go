package config

import (
	"fmt"
	"log"
	"os"

	"github.com/lpernett/godotenv"
)

type Config struct {
	PublicHost string
	Port       string

	DbUser     string
	DbPassword string
	DbAddress  string
	DbName     string
}

var Envs = initConfig()

func initConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DbUser:     getEnv("DBUSER", "root"),
		DbPassword: getEnv("DBPASSWORD", "shockwave"),
		DbAddress:  fmt.Sprintf("%s:%s", getEnv("DBHOST", "127.0.0.1"), getEnv("DBPORT", "3306")),
		DbName:     getEnv("DBNAME", "testEcom"),
	}

}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	} else {
		return fallback

	}

}
