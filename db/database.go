package db

import(
	"gorm.io/driver/mysql"
  	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Db_connection() {
	var err error

	dsn := "root:12345@tcp(127.0.0.1:3306)/Company?charset=utf8mb4"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to DB")
	}
}





