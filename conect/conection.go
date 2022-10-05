package main

import (
	"database/sql" // Pacote Database SQL para realizar Query
	"log"

	_ "github.com/go-sql-driver/mysql" // Driver Mysql para Go
)

func conect() (bd *sql.DB) {
	driver := "mysql"
	user := ""
	password := ""
	name := ""

	bd, err := sql.Open(driver, user+":"+password+"@/"+name)

	if err != nil {
		log.Print("error")
	}
	return bd
}
