package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rinha/database"
	"rinha/models"
	"rinha/repository"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

// https://raw.githubusercontent.com/zanfranceschi/rinha-de-backend-2023-q3/main/stress-test/user-files/resources/pessoas-payloads.tsv

func personRouter(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		term := r.URL.Query().Get("t")
		if term != "" {
			searchPerson(w, r)
		} else {
			getPerson(w, r)
		}
		return
	}

	if r.Method == http.MethodPost {
		createPerson(w, r)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

// GET /pessoas/[:id] – para consultar um recurso criado com a requisição anterior.
func getPerson(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/pessoas/")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	person := repository.GetPerson(id)

	err = json.NewEncoder(w).Encode(person)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}

// GET /pessoas?t=[:termo da busca] – para fazer uma busca por pessoas.
func searchPerson(w http.ResponseWriter, r *http.Request) {
	term := r.URL.Query().Get("t")
	if term == "" {
		http.Error(w, "Invalid search term", http.StatusBadRequest)
		return
	}

	// search for the term in the database
	people := repository.SearchPeople(term)

	// return the results as a JSON array
	err := json.NewEncoder(w).Encode(people)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}

// POST /pessoas – para criar um recurso pessoa.
//
//	{
//		"apelido" : "josé",
//		"nome" : "José Roberto",
//		"nascimento" : "2000-10-01",
//		"stack" : ["C#", "Node", "Oracle"]
//	}
func createPerson(w http.ResponseWriter, r *http.Request) {
	// validate the request body with models.Person
	var person models.Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	person.ID = uuid.New()

	// validate and return 422 if invalid
	if !person.Validate() {
		fmt.Println("Invalid person:", person)
		http.Error(w, "Invalid person", http.StatusUnprocessableEntity)
		return
	}

	// save the person in the database
	repository.CreatePerson(person)

	// returns status code 201 with Location header set to /pessoas/[:id]
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Location", "/pessoas/"+person.ID.String())
}

// GET /contagem-pessoas – endpoint especial para contagem de pessoas cadastradas.
func getPersonCount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	count := repository.CountPeople()
	w.Write([]byte(strconv.FormatInt(count, 10)))
}

func main() {
	database.Init()
	http.HandleFunc("/pessoas", personRouter)
	http.HandleFunc("/pessoas/", personRouter)
	http.HandleFunc("/contagem-pessoas/", getPersonCount)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}
