package main

import (
	"database/sql"
	"fmt"

	"github.com/lareza-farhan-wanaghi/gophercises/phone"
	_ "github.com/lib/pq"
)

// main provides the entry point of the app
func main() {
	db, err := sql.Open("postgres", "user=postgres dbname=Custom sslmode=disable password=")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	phoneDb, err := phone.NewPhoneDatabase(db)
	if err != nil {
		panic(err)
	}

	phoneDb.FlushDb()
	if err != nil {
		panic(err)
	}

	mockData := []phone.PhoneModel{
		{PhoneNumber: "1234567890"},
		{PhoneNumber: "123 456 7891"},
		{PhoneNumber: "(123) 456 7892"},
		{PhoneNumber: "(123) 456-7893"},
		{PhoneNumber: "123-456-7894"},
		{PhoneNumber: "123-456-7890"},
		{PhoneNumber: "1234567892"},
		{PhoneNumber: "(123)456-7892"},
	}
	fmt.Printf("Initial phone numbers:\n")
	for i, p := range mockData {
		_, err := phoneDb.InsertPhone(&p)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%d. %s\n", i+1, p.PhoneNumber)
	}

	err = phoneDb.NormalizePhoneData()
	if err != nil {
		panic(err)
	}

	phones, err := phoneDb.GetAllPhones()
	if err != nil {
		panic(err)
	}

	fmt.Printf("After normalization phone numbers:\n")
	for i, p := range phones {
		fmt.Printf("%d. %s\n", i+1, p.PhoneNumber)
	}
}
