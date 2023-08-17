package database

import (
	"fmt"
	"rinha/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func Conn() *gorm.DB {
	if database != nil {
		return database
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai", // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true,                                                                                                              // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	database = db
	return database
}

func Init() {
	CreateDatabaseExtentions()
	Migrate()
	// CreateDatabase()
	// CreatePersonTable()
}

func Migrate() {
	Conn().AutoMigrate(&models.Person{})
}

func CreateDatabaseExtentions() {
	db := Conn()
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
}

func CreateDatabase() {
	// fmt.Println("Creating database...")
	// db := Conn()
	// result := db.Exec("CREATE DATABASE IF NOT EXISTS postgres")
	// fmt.Println("CreateDatabase result: ", result)
}

func CreatePersonTable() {
	fmt.Println("Creating person table...")
	db := Conn()

	result := db.Exec(`CREATE TABLE IF NOT EXISTS person (
		id SERIAL PRIMARY KEY,
		nickname VARCHAR(32) NOT NULL,
		name VARCHAR(100) NOT NULL,
		birthday DATE NOT NULL,
		stack VARCHAR(32)[] NOT NULL
	)`)
	fmt.Println("CreatePersonTable result: ", result)
}
