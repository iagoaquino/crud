package conect

import (
	"database/sql" // Pacote Database SQL para realizar Query
	"log"

	_ "github.com/go-sql-driver/mysql" // Driver Mysql para Go
)

func Conect() (bd *sql.DB) {
	driver := "mysql"
	user := "root"
	password := "teste123"
	name := "alunos"

	bd, err := sql.Open(driver, user+":"+password+"@(172.17.0.2:3306)/"+name)

	if err != nil {
		log.Print("error com a conexão do banco")
	}
	return bd
}
