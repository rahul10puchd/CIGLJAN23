package db

import (
	"gin/basic/model"

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
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(model.Post{}) //helping to create table automatically
	db.AutoMigrate(model.User{})
	PostgresDB = db
	return db, err
}
