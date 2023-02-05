package db

import (
	"fmt"
	"gin/basic/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// func GetPostgresDBConnection() (db *gorm.DB, err error) {
// 	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
// 	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	return
// }
var PostgresDB *gorm.DB

func GetDBConn() *gorm.DB {
	return PostgresDB
}
func GetPostgresDBConnection() (*gorm.DB, error) {
	// "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DBHost"),
		os.Getenv("DBUser"),
		os.Getenv("DBPassword"),
		os.Getenv("DBName"),
		os.Getenv("DBPort"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(model.Post{}) //helping to create table automatically
	db.AutoMigrate(model.User{})
	PostgresDB = db
	return db, err
}
