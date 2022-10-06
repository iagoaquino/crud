package conect

import (
	"database/sql" // Pacote Database SQL para realizar Query
	"log"

	_ "github.com/go-sql-driver/mysql" // Driver Mysql para Go
)

func Conect() (bd *sql.DB) {
	driver := "mysql"
	user := "root"
	password := ""
	name := "alunos"

	bd, err := sql.Open(driver, user+":"+password+"@/"+name)

	if err != nil {
		log.Print("error com a conexão do banco")
	}
	return bd
}
