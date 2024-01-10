package database

import (
	"blogApp/api/initializers"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GlobalDB *gorm.DB

func ConnectDB(config *initializers.Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable ", config.DbHost, config.DbUsername, config.DbPassword, config.Dbname, config.DbPort)
	GlobalDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect ddatabase")

	}

	fmt.Println("Connected database")

}

func CloseDB() {

	db, err := GlobalDB.DB()
	if err != nil {
		log.Fatal("db couldnt close", err)
	}
	defer db.Close()
}
