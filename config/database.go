package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB
var err error

func DbInit() {
	// config := GetConfig()
	connectString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"localhost", "postgres", "822111", "book", "5432")

	// Open the connection to the database
	db, err = gorm.Open(postgres.Open(connectString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatalf("DB Connection Error: %v", err)
	}

	fmt.Println("Connected to Database")

	// Uncomment if you have models to migrate
	// err = db.AutoMigrate(&models.User{}, &models.BookSummary{})
	// if err != nil {
	// 	log.Fatalf("Migration failed: %v", err)
	// }
	// fmt.Println(config.DbHost, config.DbName)
}

func DbManager() *gorm.DB {
	return db
}
