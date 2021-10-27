package database

import (
	"fmt"
	"myapp/patients"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	DBConn *gorm.DB
)

func InitDataBase() {
	var err error

	DBConn, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to Connect DataBase")
	}
	defer DBConn.Close()
	fmt.Println("DataBase Connected Successfully")

	DBConn.AutoMigrate(&patients.Medicine{})
	fmt.Println("DataBase Migrated")
}
