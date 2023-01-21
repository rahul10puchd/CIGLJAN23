package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// func GetPostgresDBConnection() (db *gorm.DB, err error) {
// 	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
// 	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	return
// }
var PostgresDB *gorm.DB

func init() {

}
func init() {

}
func init() {

}
func init() {

}
func init() {

}
func GetPostgresDBConnection() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
