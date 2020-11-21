package DButil

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var user string = os.Getenv("DB_USER")
var password string = os.Getenv("DB_PASSWORD")
var dbname string = os.Getenv("DB_NAME")
var config string = fmt.Sprintf("%s:%s@(mysql:3306)/%s?parseTime=true",
	user,
	password,
	dbname,
)

var db, _ = gorm.Open("mysql", config)

func GetClient() *gorm.DB {
	return db
}
