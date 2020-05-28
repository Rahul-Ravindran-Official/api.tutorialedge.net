package database

import (
	"fmt"
	"os"

	"github.com/TutorialEdge/api.tutorialedge.net/challenge"
	"github.com/TutorialEdge/api.tutorialedge.net/comments"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// GetDBConn returns a pointer to a database connection
// calling functions need to ensure they defer the closing
// of this connection
func GetDBConn() (*gorm.DB, error) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := 25060

	postgresConn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", dbHost, dbPort, dbUsername, dbTable, dbPassword)

	// dbConnectionString := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + strconv.Itoa(dbPort) + ")/" + dbTable

	db, err := gorm.Open("postgres", postgresConn)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&comments.Comment{})
	db.AutoMigrate(&challenge.Challenge{})

	return db, nil
}

// Migrate migrates the database with any changes made
// the
func Migrate() {
	db, err := GetDBConn()
	if err != nil {
		fmt.Println("Could not migrate database...")
		fmt.Println(err)
	}
	db.AutoMigrate(&comments.Comment{})
}
