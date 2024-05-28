package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/lpernett/godotenv"
)

type Config struct {
	PublicHost string
	Port       string

	DbUser                 string
	DbPassword             string
	DbAddress              string
	DbName                 string
	JwtExpirationInSeconds int64
	JWTSecret              string
}

var Envs = initConfig()

func initConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return Config{
		PublicHost:             getEnv("PUBLIC_HOST", "http://localhost"),
		Port:                   getEnv("PORT", "8000"),
		DbUser:                 getEnv("DBUSER", "root"),
		DbPassword:             getEnv("DBPASSWORD", "shockwave"),
		DbAddress:              fmt.Sprintf("%s:%s", getEnv("DBHOST", "127.0.0.1"), getEnv("DBPORT", "3306")),
		DbName:                 getEnv("DBNAME", "testEcom"),
		JwtExpirationInSeconds: getEnvAsInt("JWT_EXP", 3600*24*7),
		JWTSecret:              getEnv("JWT_SECRET", "secret-no-secret-one"),
	}

}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback

}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	} else {
		return fallback

	}

}
