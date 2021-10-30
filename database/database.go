package database

import (
	"fmt"
	"myapp/organization"
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

	DBConn.LogMode(true)
	DBConn.AutoMigrate(
		&patients.Medicine{},
		&patients.Patient{},
		&organization.Organization{},
	)
	fmt.Println("DataBase Migrated")
}
