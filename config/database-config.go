package config

import (
	"fmt"
	"golang_api_hupiutang/entity"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// SetupDatabaseConnection
func SetupDatabaseConnection() *gorm.DB {
	// errEnv := godotenv.Load()

	// if errEnv != nil {
	// 	panic("Failed to load env file ")
	// }

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=require TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbPort, dbName)
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to create a connection to database ")
	}

	// nanti kita isi modelnya disini
	db.AutoMigrate(&entity.Hutang{}, &entity.User{})

	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection to database")
	}

	dbSQL.Close()
}
