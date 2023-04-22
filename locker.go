package main

import (
	"locker/cmd"
	"locker/db"
	"log"
)

func main() {
	connection, err := db.NewDB("./password.txt")
	if err != nil {
		log.Printf("error in connect to file database: %v\n", err.Error())
	}

	db.DBConnection = *connection

	cmd.GenerateCmd()
}
