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
		DSN:                  "host=db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable", // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true,                                                                                // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetMaxOpenConns(50)

	database = db
	return database
}

func Init() {
	dropAllTables()
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

func dropAllTables() {
	fmt.Println("Dropping all tables...")
	db := Conn()
	result := db.Exec(`DROP TABLE IF EXISTS people`)
	fmt.Println("dropAllTables result: ", result)
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
