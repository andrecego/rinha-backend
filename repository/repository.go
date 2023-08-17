package repository

import (
	"fmt"
	"log"
	"rinha/database"
	"rinha/models"

	"github.com/google/uuid"
)

func CreatePerson(person models.Person) {
	// save the person in the database
	db := database.Conn()
	person.ID = uuid.New()
	result := db.Create(&person)
	fmt.Println("[CreatePerson] Result: ", result)
}

// Create a func to count the number of people in the database
func CountPeople() int64 {
	db := database.Conn()
	var count int64

	// count the number of people in the database
	result := db.Model(&models.Person{}).Count(&count)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return count
}

func GetPerson(id uuid.UUID) models.Person {
	db := database.Conn()
	var person models.Person

	// find the person in the database
	result := db.First(&person, id)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return person
}

func SearchPeople(query string) []models.Person {
	db := database.Conn()
	var people []models.Person

	result := db.Where("nickname ILIKE ? OR name ILIKE ? OR ? = ANY(stack)", "%"+query+"%", "%"+query+"%", query).Find(&people)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return people
}
