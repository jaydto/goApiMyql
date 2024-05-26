package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/jaydto/goApiMyql/cmd/api"
	"github.com/jaydto/goApiMyql/config"
	"github.com/jaydto/goApiMyql/db"
)

func main() {
	db, err := db.NewMysqlStorage(mysql.Config{
		User:                 config.Envs.DbUser,
		Passwd:               config.Envs.DbPassword,
		Addr:                 config.Envs.DbAddress,
		DBName:               config.Envs.DbName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}
	initStorage(db)

	server := api.NewApiServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)

	}

}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB: Successfully Connected!")

}
