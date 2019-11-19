package configs

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	GetDB *gorm.DB // connection about DB
)

func Connect() {
	// declare variable to connect db from ENV
	host := os.Getenv("CONFIGDB_HOST")
	port := os.Getenv("CONFIGDB_PORT")
	user := os.Getenv("CONFIGDB_USER")
	password := os.Getenv("CONFIGDB_PASSWORD")
	dbname := os.Getenv("CONFIGDB_DBNAME")

	// connect database in here
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbname)
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		panic(err)
	}
	GetDB = db
}
