package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	config "github.com/magafagafa/go-app/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *sql.DB

var err error

func ConnectDb() (db *gorm.DB, err error) {
	connectTemplate := "%s:%s@%s/%s"
	connect := fmt.Sprintf(connectTemplate, config.Config.User, config.Config.Password, config.Config.DBPort, config.Config.Database) + "?charset=utf8mb4&parseTime=true&loc=Asia%2FTokyo"
	db, err = gorm.Open(mysql.Open(connect), &gorm.Config{})

	if err != nil {
		log.Println(err.Error())
	}

	return db, err
}
