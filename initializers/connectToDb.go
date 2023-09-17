package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Declare DB as a global variable so that it can be reached
// anywhere in the project
var DB *gorm.DB

func ConnectToDb() {
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
