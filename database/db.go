package database

import (
	"fmt"
	"os"
	"strconv"

	"github.com/elliotforbes/api.tutorialedge.net/comments"
	"github.com/jinzhu/gorm"
)

func GetDBConn() (*gorm.DB, error) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := 25060
	dbConnectionString := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + strconv.Itoa(dbPort) + ")/" + dbTable

	db, err := gorm.Open("mysql", dbConnectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Migrate() {
	db, err := GetDBConn()
	if err != nil {
		fmt.Println("Could not migrate database...")
		fmt.Println(err)
	}
	db.AutoMigrate(&comments.Comment{})
}
