package main

import (
	"log"
	"os"

	mysqlConfig "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jaydto/goApiMyql/config"
	"github.com/jaydto/goApiMyql/db"
)




func main(){

	db, err:=db.NewMysqlStorage(mysqlConfig.Config{
		User: config.Envs.DbUser,
		Passwd: config.Envs.DbPassword,
		Addr:config.Envs.DbAddress,
		DBName: config.Envs.DbName,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	})
	if err!=nil{
		log.Fatal(err)
	}
	driver,err:=mysql.WithInstance(db, &mysql.Config{})
	if err!=nil{
		log.Fatal(err)
	}
	m, err:=migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mysql",
		driver,
	)
	cmd:=os.Args[(len(os.Args)-1)]
	if cmd=="up"{
		if err:=m.Up();err!=nil&&err!=migrate.ErrNoChange{
			log.Fatal(err)
		}
	}
	if cmd=="down"{
		if err:=m.Down();err!=nil&&err!=migrate.ErrNoChange{
			log.Fatal(err)
		}
	}
}