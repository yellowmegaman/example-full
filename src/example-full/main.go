package main

import (
	"fmt"
	"os"
	"database/sql"
	_ "github.com/lib/pq"
)

var (
	host     = os.Getenv("PGHOST")
	user     = os.Getenv("PGUSER")
	password = os.Getenv("PGPASS")
	dbname   = os.Getenv("PGDBNAME")
)


func main() {
	connectdb() // let's start with that, since if db isn't available, why bother serving http endpoint
	apiserve()
}

func connectdb() {
	pgconnection := fmt.Sprintf("host=%s port=%d user=%s "+ "password=%s dbname=%s sslmode=disable", host, 5432, user, password, dbname)
	db, err := sql.Open("postgres", pgconnection)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to db!")
}


func apiserve() {
	
}
