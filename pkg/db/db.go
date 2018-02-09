package db

import (
	"fmt"

	"github.com/go-pg/pg"
)

// Test initializes a connection to the postgres db.
func Test() {
	db := pg.Connect(&pg.Options{
		User:     "alex",
		Password: "",
		Database: "postgres",
	})

	var n int
	_, err := db.QueryOne(pg.Scan(&n), "SELECT 1")
	if err != nil {
		panic(err)
	}
	fmt.Println(n)

	err = db.Close()
	if err != nil {
		panic(err)
	}
}
