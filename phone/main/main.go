package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	"github.com/lareza-farhan-wanaghi/gophercises/phone"
	_ "github.com/lib/pq"
)

// main provides the entry point of the app
func main() {
	dbFlag := flag.String("d", "user=postgres dbname=Custom sslmode=disable password=", "Specifies the configuration for the database connection")
	db, err := sql.Open("postgres", *dbFlag)
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

	fmt.Printf("Initial phone numbers:\n")
	for i := 1; i < len(os.Args); i++ {
		phone := &phone.PhoneModel{PhoneNumber: os.Args[i]}
		_, err := phoneDb.InsertPhone(phone)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%d. %s\n", i+1, phone.PhoneNumber)
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
